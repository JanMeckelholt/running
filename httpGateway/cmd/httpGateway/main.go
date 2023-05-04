package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/JanMeckelholt/running/common/dependencies"
	"github.com/JanMeckelholt/running/httpGateway/service"
	"github.com/JanMeckelholt/running/httpGateway/service/auth"
	"github.com/JanMeckelholt/running/httpGateway/service/mux"
	"github.com/JanMeckelholt/running/httpGateway/service/server"

	"github.com/caarlos0/env/v7"
	log "github.com/sirupsen/logrus"
)

func main() {
	var allowOriginStr string
	srv := &service.Service{}
	err := env.Parse(&srv.Config)
	if err != nil {
		return
	}
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
	rs, err := server.NewHTTPGatewayServer(srv.Clients)

	rootMux := http.NewServeMux()
	allowOriginStr = fmt.Sprintf("http://localhost:%d", srv.Config.RunningAppPort)
	if srv.Config.IsDev {
		allowOriginStr = "*"
	}
	log.Infof("allowOrigin: %s", allowOriginStr)

	rootMux.Handle(service.LoginRoute, mux.Handler(service.LoginRoute, rs))
	rootMux.Handle(service.WebsiteRoute, mux.Handler(service.WebsiteRoute, rs))
	rootMux.Handle("/health", mux.Handler("/health", rs))
	rootMux.Handle("/athlete", mux.Handler("/athlete", rs))
	rootMux.Handle("/activities", mux.Handler("/activities", rs))
	rootMux.Handle("/athlete/create", mux.Handler("/athlete/create", rs))
	rootMux.Handle("/weeksummary", mux.Handler("/weeksummary", rs))
	rootMux.Handle("/weeksummaries", mux.Handler("/weeksummaries", rs))
	rootMux.Handle("/activitiesToDB", mux.Handler("/activitiesToDB", rs))

	handlerWithAuth := server.AuthMiddleware(rootMux)
	handlerWithCors := server.CorsMiddleware(handlerWithAuth, srv.Config)

	go func() {
		sTLS := &http.Server{
			Addr:    fmt.Sprintf(":%d", dependencies.Configs["httpGatewayTLS"].Port),
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

		log.Infof("Listening on :%d", dependencies.Configs["httpGatewayTLS"].Port)
		serveErr := sTLS.ServeTLS(lis, "httpGateway/certs/http_gateway-server-cert.pem", "httpGateway/certs/http_gateway-server-key.pem")
		defer func() {
			teardown()
		}()
		if serveErr != nil {
			log.Fatal("Runner-Server: Serving Error!")
		}
	}()

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", dependencies.Configs["httpGateway"].Port),
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

	log.Infof("Listening on :%d", dependencies.Configs["httpGateway"].Port)
	serveErr := s.Serve(lis)
	defer func() {
		teardown()
	}()
	if serveErr != nil {
		log.Fatal("Runner-Server: Serving Error!")
	}

}
