package server

import (
	"context"
	log "github.com/sirupsen/logrus"
	"io"
	"net/url"

	"github.com/JanMeckelholt/running/strava-service/service/strava"

	grpcStrava "github.com/JanMeckelholt/running/common/grpc/strava"
)

type StravaServer struct {
	stravaClient *strava.Client
	grpcStrava.UnimplementedStravaServer
}

func (s StravaServer) Get(ctx context.Context, request *grpcStrava.Request) (*grpcStrava.Response, error) {
	resp, err := s.stravaClient.GetAthlet(ctx)
	if err != nil {
		log.Errorf("Could not Get Athlet")
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return &grpcStrava.Response{
		Something: string(b),
	}, nil
}

func NewServer(stravaURL url.URL, stravaToken string) *StravaServer {
	stravaClient := strava.NewClient(stravaURL, stravaToken)
	return &StravaServer{
		stravaClient: stravaClient,
	}
}
