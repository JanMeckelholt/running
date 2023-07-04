package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	certhandling "github.com/JanMeckelholt/running/common/cert-handling"
	"github.com/JanMeckelholt/running/common/commonconf"
	"github.com/JanMeckelholt/running/common/dependencies"
	"github.com/JanMeckelholt/running/common/grpc/runner"
	"github.com/JanMeckelholt/running/common/utils"
	"github.com/JanMeckelholt/running/runner/service"
	"github.com/JanMeckelholt/running/runner/service/frontend"
	"github.com/JanMeckelholt/running/runner/service/server"
)

func main() {
	commonConf := commonconf.GPGConf{}
	err := env.Parse(&commonConf)
	if err != nil {
		return
	}

	err = utils.DecryptPGP("./volumes-data/certs/runner-server-key.pem.asc", "./secret/certs/runner-server-key.pem", commonConf.GPGPrivateKey)
	if err != nil {
		log.Errorf("Could not decrypt certs: %s", err.Error())
		return
	}
	err = godotenv.Load("./volumes-data/env/.env.docker")
	if err != nil {
		log.Errorf("Could not load env: %s", err.Error())
		return
	}
	srv := &service.Service{}
	err = env.Parse(&srv.Config)
	if err != nil {
		log.Errorf("Could not load serviceConfig: %s", err.Error())
	}
	err = srv.Clients.Dial(srv.Config)
	if err != nil {
		log.Errorf("could not Dial Clients! %s", err.Error())
	}

	go func() {
		rs, err := server.NewRunnerServer(srv.Clients)
		sTLS := &http.Server{
			Addr:    fmt.Sprintf(":%d", dependencies.Configs["runner_frontend"].Port),
			Handler: frontend.FrontEnd(0, rs),
		}
		lis, err := net.Listen("tcp", sTLS.Addr)
		if err != nil {
			return
		}

		teardown := func() {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			shutdownErr := sTLS.Shutdown(ctx)
			if shutdownErr != nil {
				log.Fatal("Runner-Server: Shutdown Error!")
			}
		}

		log.Infof("Listening on :%d", dependencies.Configs["runner_frontend"].Port)
		serveErr := sTLS.ServeTLS(lis, "volumes-data/certs/runner-server-cert.pem", "secret/certs/runner-server-key.pem")
		defer func() {
			teardown()
		}()
		if serveErr != nil {
			log.Fatal("Runner-Frontend: Serving Error!")
		}
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", dependencies.Configs["runner"].Port))
	tlsCredentials, err := certhandling.LoadTLSServerCredentials("volumes-data/certs/runner-server-cert.pem", "secret/certs/runner-server-key.pem")
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(tlsCredentials))

	teardownGrpc := grpcServer.GracefulStop
	runnerServer, err := server.NewRunnerServer(srv.Clients)
	runner.RegisterRunnerServer(grpcServer, runnerServer)

	log.Infof("listening at :%d", dependencies.Configs["runner"].Port)
	serveErr := grpcServer.Serve(lis)
	defer func() {
		teardownGrpc()
	}()
	if serveErr != nil {
		log.Fatal("Runner-Server: Serving Error!")
	}

}
