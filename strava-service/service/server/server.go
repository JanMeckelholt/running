package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"

	grpcToken "github.com/JanMeckelholt/running/common/grpc/token"
	"github.com/JanMeckelholt/running/strava-service/service/clients"
	"github.com/JanMeckelholt/running/strava-service/service/strava"

	grpcStrava "github.com/JanMeckelholt/running/common/grpc/strava"
)

type StravaServer struct {
	StravaClient *strava.Client
	Clients      clients.Clients
	grpcStrava.UnimplementedStravaServer
}

func NewStravaServer(stravaURL url.URL, clients clients.Clients) *StravaServer {
	stravaClient := strava.NewClient(stravaURL)
	return &StravaServer{
		Clients:      clients,
		StravaClient: stravaClient,
	}
}

func (s StravaServer) GetToken(ctx context.Context, id *grpcToken.ClientId) (string, error) {
	resp, err := s.Clients.TokenClient.GetClient(ctx, id)
	if err != nil {
		log.Errorf("Could not Get Athlet")
		return "", err
	}
	return resp.GetToken(), nil
}

func (s StravaServer) GetClient(ctx context.Context, id *grpcToken.ClientId) (*grpcToken.Client, error) {
	return s.Clients.TokenClient.GetClient(ctx, id)
}

func (s StravaServer) GetRunner(ctx context.Context, req *grpcStrava.RunnerRequest) (*grpcStrava.RunnerResponse, error) {
	resp, err := s.StravaClient.GetAthlet(ctx, req.GetToken())
	if err != nil {
		log.Errorf("Could not Get Athlet")
		return nil, err
	}
	if resp.StatusCode == http.StatusUnauthorized {
		client, err := s.GetClient(ctx, &grpcToken.ClientId{ClientId: req.GetClientId()})
		if err != nil {
			return nil, err
		}
		newToken, err := s.StravaClient.UseRefreshToken(ctx, client.GetClientId(), client.GetClientSecret(), client.GetRefreshToken())
		if err != nil {
			return nil, err
		}
		updateReq := grpcToken.UpdateRequest{
			ClientId: req.ClientId,
			KvPairs: []*grpcToken.KvPair{
				{
					Key:   "token",
					Value: newToken,
				},
			},
		}
		client, err = s.Clients.TokenClient.UpdateClient(ctx, &updateReq)
		if err != nil {
			return nil, err
		}
		resp, err = s.StravaClient.GetAthlet(ctx, client.GetToken())
		if err != nil {
			log.Errorf("Could not Get Athlet")
			return nil, err
		}
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

func (s StravaServer) GetActivities(ctx context.Context, req *grpcStrava.ActivityRequest) (*grpcStrava.ActivitiesResponse, error) {
	resp, err := s.StravaClient.GetActivities(ctx, req.GetToken(), req.GetSince())
	if err != nil {
		log.Errorf("Could not Get Activities")
		return nil, err
	}
	if resp.StatusCode == http.StatusUnauthorized {
		client, err := s.GetClient(ctx, &grpcToken.ClientId{ClientId: req.GetClientId()})
		if err != nil {
			log.Errorf("could not get client with id %s.", req.GetClientId())
			return nil, fmt.Errorf("could not get client with id %s", req.GetClientId())
		}
		newToken, err := s.StravaClient.UseRefreshToken(ctx, client.GetClientId(), client.GetClientSecret(), client.GetRefreshToken())
		log.Infof("newToken: %s", newToken)
		if err != nil {
			return nil, err
		}
		updateReq := grpcToken.UpdateRequest{
			ClientId: req.ClientId,
			KvPairs: []*grpcToken.KvPair{
				{
					Key:   "token",
					Value: newToken,
				},
			},
		}
		client, err = s.Clients.TokenClient.UpdateClient(ctx, &updateReq)
		if err != nil {
			return nil, err
		}
		resp, err = s.StravaClient.GetActivities(ctx, client.GetToken(), req.GetSince())
		if err != nil {
			log.Errorf("Could not Get Activities")
			return nil, err
		}
	}
	if resp.StatusCode != http.StatusOK {
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
		log.Infof("Climb: %f Date: %s", activity.GetTotalElevationGain(), activity.GetStartDate())
		aR = append(aR, &activity)
	}
	_, err = dec.Token()
	if err != nil {
		log.Errorf("Could not decode Activities-Response: %s", err.Error())
		return nil, err
	}
	return &grpcStrava.ActivitiesResponse{Activities: aR}, nil
}
