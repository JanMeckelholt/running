package clients

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	certhandling "github.com/JanMeckelholt/running/common/cert-handling"
	"github.com/JanMeckelholt/running/common/dependencies"
	"github.com/JanMeckelholt/running/common/grpc/token"
	"github.com/JanMeckelholt/running/strava-service/service/config"
)

type Clients struct {
	TokenClient token.TokenClient
}

func (c *Clients) Dial(config config.ServiceConfig) error {
	conn, err := dial("token-service", config.TokenServiceName)
	if err != nil {
		return err
	}
	c.TokenClient = token.NewTokenClient(conn)

	return nil
}

func dial(serviceName, address string) (*grpc.ClientConn, error) {
	tlsCredentials, err := certhandling.LoadTLSClientCredentials("strava-service/certs/ca-cert.pem")
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
