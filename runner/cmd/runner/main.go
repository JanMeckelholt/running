package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
	"time"

	"github.com/JanMeckelholt/running/common/grpc/dependencies"
	"github.com/JanMeckelholt/running/runner/service"
	"github.com/JanMeckelholt/running/runner/service/mux"
	"github.com/JanMeckelholt/running/runner/service/server"
)

func main() {
	srv := &service.Service{}
	err := srv.Clients.Dial()
	if err != nil {
		log.Errorf("could not Dial Clients! %s", err.Error())
	}
	rs, err := server.NewRunnerServer(srv.Clients)
	rootMux := http.NewServeMux()
	rootMux.Handle("/health", mux.Handler("/health", rs))
	rootMux.Handle("/athlet", mux.Handler("/athlet", rs))
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
