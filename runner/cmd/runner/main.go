package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
	"time"

	"github.com/JanMeckelholt/running/runner/service"
	"github.com/JanMeckelholt/running/runner/service/mux"
)

func main() {
	srv := &service.Service{}

	err := srv.Clients.Dial()
	if err != nil {
		log.Errorf("could not Dial Clients!")
	}
	rootMux := http.NewServeMux()
	rootMux.Handle("/health", mux.Handler("/health"))
	rootMux.Handle("/athlet", mux.Handler("/athlet"))
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

	fmt.Printf("Listening on localholst:80")
	serveErr := server.Serve(lis)
	defer func() {
		teardown()
	}()
	if serveErr != nil {
		log.Fatal("Runner-Server: Serving Error!")
	}

}
