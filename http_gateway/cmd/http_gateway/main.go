package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/JanMeckelholt/running/common/commonconf"
	"github.com/JanMeckelholt/running/common/dependencies"
	"github.com/JanMeckelholt/running/common/utils"
	"github.com/JanMeckelholt/running/http_gateway/service"
	"github.com/JanMeckelholt/running/http_gateway/service/auth"
	"github.com/JanMeckelholt/running/http_gateway/service/config"
	"github.com/JanMeckelholt/running/http_gateway/service/mux"
	"github.com/JanMeckelholt/running/http_gateway/service/server"
	"github.com/joho/godotenv"

	"github.com/caarlos0/env/v7"
	log "github.com/sirupsen/logrus"
)

func main() {
	var allowOriginStr string
	commonConf := commonconf.GPGConf{}
	err := env.Parse(&commonConf)
	if err != nil {
		return
	}
	err = utils.DecryptPGP("./common/.env.docker.running_app.secret.asc", "./secret/env/.env.docker.running_app.secret", commonConf.GPGPrivateKey)
	if err != nil {
		log.Errorf("Could not decrypt common-env: %s", err.Error())
		return
	}
	err = utils.DecryptPGP("./volumes-data/env/.env.docker.secret.asc", "./secret/env/.env.docker.secret", commonConf.GPGPrivateKey)
	if err != nil {
		log.Errorf("Could not decrypt env: %s", err.Error())
		return
	}
	err = utils.DecryptPGP("./volumes-data/certs/http_gateway-server-key.pem.asc", "./secret/certs/http_gateway-server-key.pem", commonConf.GPGPrivateKey)
	if err != nil {
		log.Errorf("Could not decrypt certs: %s", err.Error())
		return
	}
	err = godotenv.Load("./volumes-data/env/.env.docker", "./secret/env/.env.docker.secret", "./common/.env.docker.running_app", "./secret/env/.env.docker.running_app.secret")
	if err != nil {
		log.Errorf("Could not load env: %s", err.Error())
		return
	}
	srv := &service.Service{}
	err = env.Parse(&srv.Config)
	if err != nil {
		log.Errorf("Could not load serviceConfig: %s", err.Error())
	}
	log.Infof("RunningAppPort: %d", srv.Config.RunningAppPort)
	if len(srv.Config.JWTKey) < 8 {
		log.Fatal("JWT not valid")
	}
	auth.JwtKey = []byte(srv.Config.JWTKey)

	err = auth.UpsertRunningClients(auth.RunningAppHttpClient, srv.Config.RunningAppPassword)
	if err != nil {
		log.Errorf("Could not setup password for runningApp-Client %s", auth.RunningAppHttpClient)
	}
	err = auth.UpsertHttpClients(auth.MasterHttpClient, srv.Config.MasterPassword)
	if err != nil {
		log.Errorf("Could not setup password for HTTP-Client %s", auth.MasterHttpClient)
	}

	err = srv.Clients.Dial(srv.Config)
	if err != nil {
		log.Errorf("could not Dial Clients! %s", err.Error())
	}
	rs, err := server.NewHttpGatewayServer(srv.Clients)

	rootMux := http.NewServeMux()
	log.Infof("allowOrigin: %s", allowOriginStr)

	rootMux.Handle(service.LoginRoute, mux.Handler(service.LoginRoute, rs))
	rootMux.Handle(service.WebsiteRoute, mux.Handler(service.WebsiteRoute, rs))
	rootMux.Handle(config.Prefix+"/health", mux.Handler("/health", rs))
	rootMux.Handle(config.Prefix+"/athlete", mux.Handler("/athlete", rs))
	rootMux.Handle(config.Prefix+"/activities", mux.Handler("/activities", rs))
	rootMux.Handle(config.Prefix+"/athlete/create", mux.Handler("/athlete/create", rs))
	rootMux.Handle(config.Prefix+"/weeksummary", mux.Handler("/weeksummary", rs))
	rootMux.Handle(config.Prefix+"/weeksummaries", mux.Handler("/weeksummaries", rs))
	rootMux.Handle(config.Prefix+"/activitiesToDB", mux.Handler("/activitiesToDB", rs))

	handlerWithAuth := server.AuthMiddleware(rootMux)
	handlerWithCors := server.CorsMiddleware(handlerWithAuth, srv.Config)

	go func() {
		sTLS := &http.Server{
			Addr:    fmt.Sprintf(":%d", dependencies.Configs["http_gatewayTLS"].Port),
			Handler: handlerWithCors,
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

		log.Infof("Listening on :%d", dependencies.Configs["http_gatewayTLS"].Port)
		serveErr := sTLS.ServeTLS(lis, "volumes-data/certs/http_gateway-server-cert.pem", "secret/certs/http_gateway-server-key.pem")
		defer func() {
			teardown()
		}()
		if serveErr != nil {
			log.Fatal("Runner-Server: Serving Error!")
		}
	}()

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", dependencies.Configs["http_gateway"].Port),
		Handler: handlerWithCors,
	}
	lis, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return
	}

	teardown := func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		shutdownErr := s.Shutdown(ctx)
		if shutdownErr != nil {
			log.Fatal("Runner-Server: Shutdown Error!")
		}
	}

	log.Infof("Listening on :%d", dependencies.Configs["http_gateway"].Port)
	serveErr := s.Serve(lis)
	defer func() {
		teardown()
	}()
	if serveErr != nil {
		log.Fatal("Runner-Server: Serving Error!")
	}

}
