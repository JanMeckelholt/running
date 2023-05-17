package main

import (
	"fmt"
	"net"
	"os"

	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	certhandling "github.com/JanMeckelholt/running/common/cert-handling"
	"github.com/JanMeckelholt/running/common/dependencies"
	grpcDB "github.com/JanMeckelholt/running/common/grpc/database"

	"github.com/JanMeckelholt/running/database-service/service/server"

	"github.com/JanMeckelholt/running/database-service/service"
)

func main() {
	wd, _ := os.Getwd()
	log.Infof("current dir: %s", wd)
	godotenv.Load("./common/.env.docker.postgres", "./common/.env.docker.postgres.secret")
	log.Infof("POSTGRES_PASSWORD: %s", os.Getenv("POSTGRES_PASSWORD"))
	for k, v := range os.Environ() {
		log.Infof("variable %d: %s\n", k, v)
	}
	storer := &service.Storer{}
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
	err = storer.AutoMigrate(service.DBActivity{})
	if err != nil {
		log.Errorf("could not automigrate DBActivity %s", err.Error())
		return
	}
	err = storer.AutoMigrate(service.DBClient{})
	if err != nil {
		log.Errorf("could not automigrate DBClient %s", err.Error())
		return
	}

	dbServer := server.NewServer(storer)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", dependencies.Configs["database-service"].Port))
	if err != nil {
		log.Fatal("cannot setup tcp listener for database-service: ", err)
	}

	tlsCredentials, err := certhandling.LoadTLSServerCredentials("database-service/certs/database-service-server-cert.pem", "database-service/certs/database-service-server-key.pem")
	if err != nil {
		log.Fatal("cannot load TLS credentials for database-service: ", err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(tlsCredentials))
	teardown := grpcServer.GracefulStop

	grpcDB.RegisterDatabaseServer(grpcServer, dbServer)

	log.Infof("listening at :%d", dependencies.Configs["database-service"].Port)
	serveErr := grpcServer.Serve(lis)
	defer func() {
		teardown()
	}()
	if serveErr != nil {
		log.Fatal("Runner-Server: Serving Error!")
	}

}
