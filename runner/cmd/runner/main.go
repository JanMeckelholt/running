package main

import (
	"context"
	"github.com/caarlos0/env/v7"
	"log"
	"net"
	"net/http"
	"time"

	"running/common/health"
	"running/runner/service"
	"running/runner/service/config"
)

func main() {
	srv := &service.Service{}
	srv.ServiceConfig = config.DefaultServiceConfig()
	err := env.Parse(&srv.ServiceConfig)
	if err != nil {
		return
	}
	rootMux := http.NewServeMux()
	rootMux.Handle("/health", health.Handler(health.Health{ServiceName: "runner"}))

	server := &http.Server{
		Addr:    ":80",
		Handler: rootMux,
	}
	lis, err := net.Listen("tcp", server.Addr)
	if err != nil {
		return
	}

	teardown := func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		shutdownErr := server.Shutdown(ctx)
		if shutdownErr != nil {
			log.Fatal("Runner-Server: Shutdown Error!")
		}
	}

	serveErr := server.Serve(lis)
	defer func() {
		teardown()
	}()
	if serveErr != nil {
		log.Fatal("Runner-Server: Serving Error!")
	}

}
