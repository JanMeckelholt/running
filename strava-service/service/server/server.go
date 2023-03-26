package server

import (
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/url"

	"github.com/JanMeckelholt/running/strava-service/service/strava"

	grpcStrava "github.com/JanMeckelholt/running/common/grpc/strava"
)

type StravaServer struct {
	stravaClient *strava.Client
	grpcStrava.UnimplementedStravaServer
}

func (s StravaServer) GetRunner(ctx context.Context, req *grpcStrava.RunnerRequest) (*grpcStrava.RunnerResponse, error) {
	resp, err := s.stravaClient.GetAthlet(ctx, req.GetToken())
	if err != nil {
		log.Errorf("Could not Get Athlet")
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Errorf("could not Get Athlet - Status-Code: %d", resp.StatusCode)
		return nil, fmt.Errorf("could not Get Athlet - Status-Code: %d", resp.StatusCode)
	}
	res := &grpcStrava.RunnerResponse{}
	err = json.NewDecoder(resp.Body).Decode(res)
	log.Infof("GetRunner response: %s %s", res.GetFirstname(), res.GetLastname())
	if err != nil {
		log.Errorf("Could not Decode Runner-Response: %s", err.Error())
		return nil, err
	}
	return res, nil
}

func NewServer(stravaURL url.URL) *StravaServer {
	stravaClient := strava.NewClient(stravaURL)
	ss := new(StravaServer)
	ss.stravaClient = stravaClient
	return ss
}
