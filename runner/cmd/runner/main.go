package main

import (
	"fmt"
	"net"

	"github.com/caarlos0/env/v7"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	certhandling "github.com/JanMeckelholt/running/common/cert-handling"
	"github.com/JanMeckelholt/running/common/dependencies"
	"github.com/JanMeckelholt/running/common/grpc/runner"
	"github.com/JanMeckelholt/running/runner/service"
	"github.com/JanMeckelholt/running/runner/service/server"
)

func main() {
	srv := &service.Service{}
	err := env.Parse(&srv.Config)
	if err != nil {
		return
	}
	err = srv.Clients.Dial(srv.Config)
	if err != nil {
		log.Errorf("could not Dial Clients! %s", err.Error())
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", dependencies.Configs["runner"].Port))
	tlsCredentials, err := certhandling.LoadTLSServerCredentials("runner/certs/runner-server-cert.pem", "runner/certs/runner-server-key.pem")
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
