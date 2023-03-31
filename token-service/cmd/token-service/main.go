package main

import (
	"fmt"
	"github.com/caarlos0/env/v7"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"

	certhandling "github.com/JanMeckelholt/running/common/cert-handling"
	"github.com/JanMeckelholt/running/common/dependencies"
	grpcToken "github.com/JanMeckelholt/running/common/grpc/token"

	"github.com/JanMeckelholt/running/token-service/service/server"

	"github.com/JanMeckelholt/running/token-service/service"
)

func main() {
	storer := &service.TokenStorer{}
	err := env.Parse(&storer.StorerConfig)
	if err != nil {
		return
	}
	storer = service.NewStorer(storer.StorerConfig)
	err = storer.InitStorage()
	if err != nil {
		log.Errorf("could not init storage %s", err.Error())
		return
	}
	err = storer.AutoMigrate(service.DBClient{})
	if err != nil {
		log.Errorf("could not automigrate storage %s", err.Error())
		return
	}

	tokenServer := server.NewServer(storer)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", dependencies.Configs["token-service"].Port))

	tlsCredentials, err := certhandling.LoadTLSServerCredentials("token-service/certs/token-service-server-cert.pem", "token-service/certs/token-service-server-key.pem")
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(tlsCredentials))
	teardown := grpcServer.GracefulStop

	grpcToken.RegisterTokenServer(grpcServer, tokenServer)

	log.Infof("listening at :%d", dependencies.Configs["token-service"].Port)
	serveErr := grpcServer.Serve(lis)
	defer func() {
		teardown()
	}()
	if serveErr != nil {
		log.Fatal("Runner-Server: Serving Error!")
	}

}
