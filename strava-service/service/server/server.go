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

func (s StravaServer) GetToken(ctx context.Context, id *grpcDB.AthleteId) (string, error) {
	resp, err := s.Clients.DatabaseClient.GetAthlete(ctx, id)
	if err != nil {
		log.Errorf("Could not Get Athlet")
		return "", err
	}
	return resp.GetToken(), nil
}

func (s StravaServer) GetClient(ctx context.Context, id *grpcDB.ClientId) (*grpcDB.Client, error) {
	return s.Clients.DatabaseClient.GetClient(ctx, id)
}

func (s StravaServer) GetAthlete(ctx context.Context, req *grpcStrava.AthleteRequest) (*grpcStrava.AthleteResponse, error) {
	var (
		token string
		err   error
	)
	token = req.GetToken()
	if token == "" {
		token, err = s.getTokenForAthelte(ctx, req.AthleteId)
		if err != nil {
			return nil, err
		}
	}
	resp, err := s.StravaClient.GetAthlet(ctx, token)
	if err != nil {
		log.Errorf("Could not Get Athlet")
		return nil, err
	}
	if resp.StatusCode == http.StatusUnauthorized {
		dbAthlete, err := s.refreshToken(ctx, req.GetAthleteId())
		if err != nil {
			return nil, err
		}
		resp, err = s.StravaClient.GetAthlet(ctx, dbAthlete.GetToken())
		if err != nil {
			log.Errorf("Could not Get Athlet")
			return nil, err
		}
	}
	if resp.StatusCode != 200 {
		log.Errorf("could not Get Athlet - Status-Code: %d", resp.StatusCode)
		return nil, fmt.Errorf("could not Get Athlet - Status-Code: %d", resp.StatusCode)
	}
	res := &grpcStrava.AthleteResponse{}
	err = json.NewDecoder(resp.Body).Decode(res)
	log.Infof("GetAthlete response: %s %s", res.GetFirstname(), res.GetLastname())
	if err != nil {
		log.Errorf("Could not decode Athlete-Response: %s", err.Error())
		return nil, err
	}
	return res, nil
}

func (s StravaServer) GetActivities(ctx context.Context, req *grpcStrava.ActivitiesRequest) (*grpcStrava.ActivitiesResponse, error) {
	var (
		token string
		err   error
	)
	token = req.GetToken()
	if token == "" {
		token, err = s.getTokenForAthelte(ctx, req.AthleteId)
		if err != nil {
			return nil, err
		}
	}
	resp, err := s.StravaClient.GetActivities(ctx, token, req.GetSince())
	if err != nil {
		log.Errorf("Could not Get Activities")
		return nil, err
	}
	if resp.StatusCode == http.StatusUnauthorized {
		dbAthlete, err := s.refreshToken(ctx, req.GetAthleteId())
		if err != nil {
			return nil, err
		}
		resp, err = s.StravaClient.GetActivities(ctx, dbAthlete.GetToken(), req.GetSince())
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

func (s StravaServer) getTokenForAthelte(ctx context.Context, athleteId uint64) (string, error) {
	dbAthlete, err := s.Clients.DatabaseClient.GetAthlete(ctx, &grpcDB.AthleteId{AthleteId: athleteId})
	if err != nil {
		log.Errorf("could not get client with id %d.", athleteId)
		return "", fmt.Errorf("could not get client with id %d", athleteId)
	}
	return dbAthlete.GetToken(), nil
}

func (s StravaServer) refreshToken(ctx context.Context, athleteId uint64) (*grpcDB.Athlete, error) {
	athlete, err := s.Clients.DatabaseClient.GetAthlete(ctx, &grpcDB.AthleteId{AthleteId: athleteId})
	if err != nil {
		log.Errorf("could not get athlete with id %s.", athleteId)
		return nil, fmt.Errorf("could not get athlere with id %s", athleteId)
	}
	client, err := s.GetClient(ctx, &grpcDB.ClientId{ClientId: athlete.ClientId})
	if err != nil {
		log.Errorf("could not get client with id %s.", athlete.ClientId)
		return nil, fmt.Errorf("could not get client with id %s", athlete.ClientId)
	}
	newToken, err := s.StravaClient.UseRefreshToken(ctx, client.GetClientId(), client.GetClientSecret(), athlete.GetRefreshToken())
	log.Infof("newToken: %s", newToken)
	if err != nil {
		return nil, err
	}
	updateReq := grpcDB.UpdateAthleteRequest{
		AthleteId: athleteId,
		KvPairs: []*grpcDB.KvPair{
			{
				Key:   "token",
				Value: newToken,
			},
		},
	}
	athlete, err = s.Clients.DatabaseClient.UpdateAthlete(ctx, &updateReq)
	if err != nil {
		return nil, err
	}
	return athlete, nil
}
