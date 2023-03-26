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
		log.Errorf("Could not decode Runner-Response: %s", err.Error())
		return nil, err
	}
	return res, nil
}

func (s StravaServer) GetActivities(ctx context.Context, req *grpcStrava.RunnerRequest) (*grpcStrava.ActivitiesResponse, error) {
	resp, err := s.stravaClient.GetActivities(ctx, req.GetToken())
	if err != nil {
		log.Errorf("Could not Get Activities")
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Errorf("could not Get Activities - Status-Code: %d", resp.StatusCode)
		return nil, fmt.Errorf("could not Get Activities - Status-Code: %d", resp.StatusCode)
	}

	var aR []*grpcStrava.Activity
	var activity grpcStrava.Activity
	dec := json.NewDecoder(resp.Body)
	_, err = dec.Token()
	if err != nil {
		log.Errorf("Could not decode Activities-Response: %s", err.Error())
		return nil, err
	}
	for dec.More() {
		err = dec.Decode(&activity)
		if err != nil {
			log.Errorf("Could not unmarshal Activity: %s", err.Error())
			return nil, err
		}
		aR = append(aR, &activity)
	}
	_, err = dec.Token()
	if err != nil {
		log.Errorf("Could not decode Activities-Response: %s", err.Error())
		return nil, err
	}
	return &grpcStrava.ActivitiesResponse{Activities: aR}, nil
}

func NewServer(stravaURL url.URL) *StravaServer {
	stravaClient := strava.NewClient(stravaURL)
	ss := new(StravaServer)
	ss.stravaClient = stravaClient
	return ss
}
