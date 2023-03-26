package clients

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/JanMeckelholt/running/common/grpc/dependencies"
	grpcStrava "github.com/JanMeckelholt/running/common/grpc/strava"
)

type Clients struct {
	Strava grpcStrava.StravaClient
}

func (c *Clients) Dial() error {
	conn, err := dial("strava-service")
	if err != nil {
		return err
	}
	c.Strava = grpcStrava.NewStravaClient(conn)
	return nil
}

func dial(serviceName string) (*grpc.ClientConn, error) {
	log.Infof("Dialing: %s:%d", serviceName, dependencies.Configs["strava-service"].Port)
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", serviceName, dependencies.Configs["strava-service"].Port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	//log.Infof("Dialing: :%d", dependencies.Configs["strava-service"].Port)
	//conn, err := grpc.Dial(fmt.Sprintf(":%d", dependencies.Configs["strava-service"].Port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	log.Infof("Dialed  %s", conn.Target())
	if err != nil {
		return nil, err
	}
	return conn, nil
}
