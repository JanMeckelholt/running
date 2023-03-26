package main

import (
	"fmt"
	"github.com/caarlos0/env/v7"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"

	"github.com/JanMeckelholt/running/common/grpc/dependencies"
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
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", dependencies.Configs["strava-service"].Port))
	grpcServer := grpc.NewServer()

	teardown := grpcServer.GracefulStop

	log.Infof("listening at :%d", dependencies.Configs["strava-service"].Port)
	serveErr := grpcServer.Serve(lis)
	defer func() {
		teardown()
	}()
	if serveErr != nil {
		log.Fatal("Runner-Server: Serving Error!")
	}

	stravaServer := server.NewServer(srv.ServiceConfig.StravaURL, srv.ServiceConfig.StravaToken)
	strava.RegisterStravaServer(grpcServer, stravaServer)
}
