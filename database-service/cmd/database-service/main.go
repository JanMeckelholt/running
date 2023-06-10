package main

import (
	"fmt"
	"net"

	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	certhandling "github.com/JanMeckelholt/running/common/cert-handling"
	"github.com/JanMeckelholt/running/common/commonconf"
	"github.com/JanMeckelholt/running/common/dependencies"
	grpcDB "github.com/JanMeckelholt/running/common/grpc/database"
	"github.com/JanMeckelholt/running/common/utils"

	"github.com/JanMeckelholt/running/database-service/service/server"

	"github.com/JanMeckelholt/running/database-service/service"
)

func main() {
	commonConf := commonconf.GPGConf{}
	err := env.Parse(&commonConf)
	if err != nil {
		log.Errorf("Could not load commonConf: %s", err.Error())
		return
	}
	err = utils.DecryptPGP("./common/.env.docker.postgres.secret.asc", "./secret/env/.env.docker.postgres.secret", commonConf.GPGPrivateKey)
	if err != nil {
		log.Errorf("Could not decrypt postgres.secret: %s", err.Error())
		return
	}
	err = utils.DecryptPGP("./volumes-data/certs/database-service-server-key.pem.asc", "./secret/certs/database-service-server-key.pem", commonConf.GPGPrivateKey)
	if err != nil {
		log.Errorf("Could not decrypt database-service-server-key.pem: %s", err.Error())
		return
	}
	err = utils.DecryptPGP("./volumes-data/certs/postgres-key.pem.asc", "./secret/certs/postgres-key.pem", commonConf.GPGPrivateKey)
	if err != nil {
		log.Errorf("Could not decrypt postgres-key.pem: %s", err.Error())
		return
	}
	err = godotenv.Load("./common/.env.docker.postgres", "./secret/env/.env.docker.postgres.secret")
	if err != nil {
		log.Errorf("Could not load env: %s", err.Error())
		return
	}

	storer := &service.Storer{}
	err = env.Parse(&storer.StorerConfig)
	if err != nil {
		log.Errorf("Could not load storerConfig: %s", err.Error())
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

	tlsCredentials, err := certhandling.LoadTLSServerCredentials("volumes-data/certs/database-service-server-cert.pem", "secret/certs/database-service-server-key.pem")
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
