package main

import (
	"context"
	"fmt"
	"github.com/caarlos0/env/v7"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
	"time"

	"github.com/JanMeckelholt/running/common/dependencies"
	"github.com/JanMeckelholt/running/runner/service"
	"github.com/JanMeckelholt/running/runner/service/mux"
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
	rs, err := server.NewRunnerServer(srv.Clients)
	rootMux := http.NewServeMux()
	rootMux.Handle("/health", mux.Handler("/health", rs))
	rootMux.Handle("/athlete", mux.Handler("/athlete", rs))
	rootMux.Handle("/activities", mux.Handler("/activities", rs))
	rootMux.Handle("/athlete/create", mux.Handler("/athlete/create", rs))
	rootMux.Handle("/weeklyclimb", mux.Handler("/weeklyclimb", rs))
	rootMux.Handle("/stravaActivitiesToDB", mux.Handler("/stravaActivitiesToDB", rs))
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", dependencies.Configs["runner"].Port),
		Handler: rootMux,
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

	log.Infof("Listening on :%d", dependencies.Configs["runner"].Port)
	serveErr := s.Serve(lis)
	defer func() {
		teardown()
	}()
	if serveErr != nil {
		log.Fatal("Runner-Server: Serving Error!")
	}

}
