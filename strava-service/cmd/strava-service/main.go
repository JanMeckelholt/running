package main

import (
	"fmt"
	"net"
	"net/url"

	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	certhandling "github.com/JanMeckelholt/running/common/cert-handling"
	"github.com/JanMeckelholt/running/common/commonconf"
	"github.com/JanMeckelholt/running/common/dependencies"
	"github.com/JanMeckelholt/running/common/grpc/strava"
	"github.com/JanMeckelholt/running/common/utils"
	"github.com/JanMeckelholt/running/strava-service/service"
	"github.com/JanMeckelholt/running/strava-service/service/server"
)

func main() {
	gpgConf := commonconf.GPGConf{}
	err := env.Parse(&gpgConf)
	if err != nil {
		log.Errorf("Could not load commonConf: %s", err.Error())
		return
	}
	err = utils.DecryptPGP("./volumes-data/certs/strava-service-server-key.pem.asc", "./secret/certs/strava-service-server-key.pem", gpgConf.GPGPrivateKey)
	if err != nil {
		log.Errorf("Could not load DecryptPGP: %s", err.Error())
		return
	}

	err = godotenv.Load("./volumes-data/env/.env.docker")
	if err != nil {
		log.Errorf("Could not load env: %s", err.Error())
		return
	}
	srv := service.Service{}
	err = env.Parse(&srv.ServiceConfig)
	if err != nil {
		log.Errorf("Could not load serviceConfig: %s", err.Error())
	}
	err = srv.Clients.Dial(srv.ServiceConfig)
	if err != nil {
		log.Errorf("could not Dial Clients! %s", err.Error())
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", dependencies.Configs["strava-service"].Port))
	tlsCredentials, err := certhandling.LoadTLSServerCredentials("volumes-data/certs/strava-service-server-cert.pem", "secret/certs/strava-service-server-key.pem")
	if err != nil {
		log.Errorf("cannot load TLS credentials: %s", err.Error())
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
