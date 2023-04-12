package main

import (
	"fmt"
	"net"
	"net/url"

	"github.com/caarlos0/env/v7"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	certhandling "github.com/JanMeckelholt/running/common/cert-handling"
	"github.com/JanMeckelholt/running/common/dependencies"
	"github.com/JanMeckelholt/running/common/grpc/strava"
	"github.com/JanMeckelholt/running/strava-service/service"
	"github.com/JanMeckelholt/running/strava-service/service/server"
)

func main() {
	srv := &service.Service{}
	err := env.Parse(&srv.ServiceConfig)
	if err != nil {
		return
	}
	err = srv.Clients.Dial(srv.ServiceConfig)
	if err != nil {
		log.Errorf("could not Dial Clients! %s", err.Error())
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", dependencies.Configs["strava-service"].Port))
	tlsCredentials, err := certhandling.LoadTLSServerCredentials("strava-service/certs/strava-service-server-cert.pem", "strava-service/certs/strava-service-server-key.pem")
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(tlsCredentials))

	teardown := grpcServer.GracefulStop
	stravaServer := server.NewStravaServer(url.URL{Host: srv.ServiceConfig.StravaURL}, srv.Clients)
	strava.RegisterStravaServer(grpcServer, stravaServer)

	log.Infof("listening at :%d", dependencies.Configs["strava-service"].Port)
	serveErr := grpcServer.Serve(lis)
	defer func() {
		teardown()
	}()
	if serveErr != nil {
		log.Fatal("Strava-Server: Serving Error!")
	}

}
