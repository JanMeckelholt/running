package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"

	grpcDB "github.com/JanMeckelholt/running/common/grpc/database"
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

func (s StravaServer) GetToken(ctx context.Context, id *grpcDB.ClientId) (string, error) {
	resp, err := s.Clients.DatabaseClient.GetClient(ctx, id)
	if err != nil {
		log.Errorf("Could not Get Athlet")
		return "", err
	}
	return resp.GetToken(), nil
}

func (s StravaServer) GetClient(ctx context.Context, id *grpcDB.ClientId) (*grpcDB.Client, error) {
	return s.Clients.DatabaseClient.GetClient(ctx, id)
}

func (s StravaServer) GetRunner(ctx context.Context, req *grpcStrava.RunnerRequest) (*grpcStrava.RunnerResponse, error) {
	token := req.GetToken()
	if token == "" {
		client, err := s.GetClient(ctx, &grpcDB.ClientId{ClientId: req.GetClientId()})
		if err != nil {
			log.Errorf("could not get client with id %s.", req.GetClientId())
			return nil, fmt.Errorf("could not get client with id %s", req.GetClientId())
		}
		token = client.GetToken()
	}
	resp, err := s.StravaClient.GetAthlet(ctx, token)
	if err != nil {
		log.Errorf("Could not Get Athlet")
		return nil, err
	}
	if resp.StatusCode == http.StatusUnauthorized {
		client, err := s.GetClient(ctx, &grpcDB.ClientId{ClientId: req.GetClientId()})
		if err != nil {
			return nil, err
		}
		newToken, err := s.StravaClient.UseRefreshToken(ctx, client.GetClientId(), client.GetClientSecret(), client.GetRefreshToken())
		if err != nil {
			return nil, err
		}
		updateReq := grpcDB.UpdateRequest{
			ClientId: req.ClientId,
			KvPairs: []*grpcDB.KvPair{
				{
					Key:   "token",
					Value: newToken,
				},
			},
		}
		client, err = s.Clients.DatabaseClient.UpdateClient(ctx, &updateReq)
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

func (s StravaServer) GetActivities(ctx context.Context, req *grpcStrava.ActivitiesRequest) (*grpcStrava.ActivitiesResponse, error) {
	token := req.GetToken()
	if token == "" {
		client, err := s.GetClient(ctx, &grpcDB.ClientId{ClientId: req.GetClientId()})
		if err != nil {
			log.Errorf("could not get client with id %s.", req.GetClientId())
			return nil, fmt.Errorf("could not get client with id %s", req.GetClientId())
		}
		token = client.GetToken()
	}
	resp, err := s.StravaClient.GetActivities(ctx, token, req.GetSince())
	if err != nil {
		log.Errorf("Could not Get Activities")
		return nil, err
	}
	if resp.StatusCode == http.StatusUnauthorized {
		client, err := s.GetClient(ctx, &grpcDB.ClientId{ClientId: req.GetClientId()})
		if err != nil {
			log.Errorf("could not get client with id %s.", req.GetClientId())
			return nil, fmt.Errorf("could not get client with id %s", req.GetClientId())
		}
		newToken, err := s.StravaClient.UseRefreshToken(ctx, client.GetClientId(), client.GetClientSecret(), client.GetRefreshToken())
		log.Infof("newToken: %s", newToken)
		if err != nil {
			return nil, err
		}
		updateReq := grpcDB.UpdateRequest{
			ClientId: req.ClientId,
			KvPairs: []*grpcDB.KvPair{
				{
					Key:   "token",
					Value: newToken,
				},
			},
		}
		client, err = s.Clients.DatabaseClient.UpdateClient(ctx, &updateReq)
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
	dec := json.NewDecoder(resp.Body)
	_, err = dec.Token()
	if err != nil {
		log.Errorf("Could not decode Activities-Response: %s", err.Error())
		return nil, err
	}
	for dec.More() {
		var activity grpcStrava.Activity
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

func (s StravaServer) ActivitiesToDB(ctx context.Context, req *grpcStrava.ActivitiesRequest) (*emptypb.Empty, error) {
	activites, err := s.GetActivities(ctx, req)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	if len(activites.GetActivities()) == 0 {
		return &emptypb.Empty{}, nil
	}
	for _, activity := range activites.GetActivities() {
		_, err = s.Clients.DatabaseClient.UpsertActivity(ctx, activity)
		if err != nil {
			return &emptypb.Empty{}, err
		}
	}
	return &emptypb.Empty{}, nil
}
