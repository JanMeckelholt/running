package clients

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/JanMeckelholt/running/common/grpc/dependencies"
	grpcStrava "github.com/JanMeckelholt/running/common/grpc/strava"
)

type Clients struct {
	Strava grpcStrava.StravaClient
}

func (c *Clients) Dial() error {
	log.Infof("Dialing: %s", dependencies.Configs["strava-service"].Address)
	conn, err := grpc.Dial(dependencies.Configs["strava-service"].Address)
	log.Infof("Dialed strava-service %s", conn)
	if err != nil {
		return err
	}
	c.Strava = grpcStrava.NewStravaClient(conn)
	return nil
}
