package clients

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	certhandling "github.com/JanMeckelholt/running/common/cert-handling"
	"github.com/JanMeckelholt/running/common/dependencies"
	"github.com/JanMeckelholt/running/common/grpc/database"
	grpcStrava "github.com/JanMeckelholt/running/common/grpc/strava"
	"github.com/JanMeckelholt/running/runner/service/config"
)

type Clients struct {
	StravaClient   grpcStrava.StravaClient
	DatabaseClient database.DatabaseClient
}

func (c *Clients) Dial(config config.ServiceConfig) error {
	conn, err := dial("strava-service", config.StravaServiceName)
	if err != nil {
		return err
	}
	c.StravaClient = grpcStrava.NewStravaClient(conn)

	conn, err = dial("database-service", config.DatabaseServiceName)
	if err != nil {
		return err
	}
	c.DatabaseClient = database.NewDatabaseClient(conn)

	return nil
}

func dial(serviceName, address string) (*grpc.ClientConn, error) {
	tlsCredentials, err := certhandling.LoadTLSClientCredentials("runner/certs/ca-cert.pem")
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}
	log.Infof("Dialing: %s:%d", address, dependencies.Configs[serviceName].Port)
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", address, dependencies.Configs[serviceName].Port), grpc.WithTransportCredentials(tlsCredentials))
	log.Infof("Dialed  %s", conn.Target())
	if err != nil {
		return nil, err
	}
	return conn, nil
}
