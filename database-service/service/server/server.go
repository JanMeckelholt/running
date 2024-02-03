package server

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/JanMeckelholt/running/database-service/service"

	grpcDB "github.com/JanMeckelholt/running/common/grpc/database"
	grpcStrava "github.com/JanMeckelholt/running/common/grpc/strava"

	log "github.com/sirupsen/logrus"
)

type DatabaseServer struct {
	Storer *service.Storer
	grpcDB.UnimplementedDatabaseServer
}

func NewServer(storer *service.Storer) *DatabaseServer {

	ts := new(DatabaseServer)
	ts.Storer = storer
	return ts
}

func (s DatabaseServer) UpsertClient(ctx context.Context, req *grpcDB.Client) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.Storer.UpsertClient(req.GetClientId(), req.GetClientSecret())
}

func (s DatabaseServer) UpdateClient(ctx context.Context, req *grpcDB.UpdateClientRequest) (*grpcDB.Client, error) {
	return s.Storer.UpdateClient(req.GetClientId(), req.GetKvPairs())
}

func (s DatabaseServer) GetClient(ctx context.Context, req *grpcDB.ClientId) (*grpcDB.Client, error) {
	dbClient, err := s.Storer.GetDBClient(req.GetClientId())
	if err != nil {
		log.Errorf("error getting dbclient :%s", err.Error())
		return nil, err
	}
	if dbClient == nil {
		log.Errorf("got nil client :%s", err.Error())
		return nil, nil
	}
	return &grpcDB.Client{
		ClientId:     *dbClient.ClientId,
		ClientSecret: *dbClient.ClientSecret,
	}, nil
}

func (s DatabaseServer) UpsertAthlete(ctx context.Context, req *grpcDB.Athlete) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.Storer.UpsertAthlete(req.GetAthleteId(), req.GetClientId(), req.GetToken(), req.GetRefreshToken())
}

func (s DatabaseServer) UpdateAthlete(ctx context.Context, req *grpcDB.UpdateAthleteRequest) (*grpcDB.Athlete, error) {
	return s.Storer.UpdateAthlete(req.GetAthleteId(), req.GetKvPairs())
}

func (s DatabaseServer) GetAthlete(ctx context.Context, req *grpcDB.AthleteId) (*grpcDB.Athlete, error) {
	dbAthlete, err := s.Storer.GetDBAthlete(req.GetAthleteId())
	if err != nil {
		log.Errorf("error getting dbathlete :%s", err.Error())
		return nil, err
	}
	if dbAthlete == nil {
		log.Errorf("got nil athlete :%s", err.Error())
		return nil, nil
	}
	return &grpcDB.Athlete{
		AthleteId:    *dbAthlete.AthleteId,
		ClientId:     *dbAthlete.ClientId,
		Token:        *dbAthlete.Token,
		RefreshToken: *dbAthlete.RefreshToken,
	}, nil
}

func (s DatabaseServer) GetActivities(ctx context.Context, req *grpcDB.ActivitiesRequest) (*grpcDB.ActivitiesResponse, error) {
	return s.Storer.GetActivities(req)
}

func (s DatabaseServer) UpsertActivityFromCSV(ctx context.Context, req *grpcStrava.Activity) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.Storer.UpsertActivity(req, true)
}
