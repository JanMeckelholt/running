package server

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/JanMeckelholt/running/database-service/service"

	grpcDB "github.com/JanMeckelholt/running/common/grpc/database"
	grpcStrava "github.com/JanMeckelholt/running/common/grpc/strava"
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
	return &emptypb.Empty{}, s.Storer.UpsertClient(req.GetClientId(), req.GetClientSecret(), req.GetToken(), req.GetRefreshToken(), uint64(req.GetAthleteId()))
}

func (s DatabaseServer) UpdateClient(ctx context.Context, req *grpcDB.UpdateRequest) (*grpcDB.Client, error) {
	return s.Storer.UpdateClient(req.GetClientId(), req.GetKvPairs())
}

func (s DatabaseServer) GetClient(ctx context.Context, req *grpcDB.ClientId) (*grpcDB.Client, error) {
	dbClient, err := s.Storer.GetDBClient(req.GetClientId())
	if err != nil {
		return nil, err
	}
	if dbClient == nil {
		return nil, nil
	}
	return &grpcDB.Client{
		ClientId:     *dbClient.ClientId,
		ClientSecret: *dbClient.ClientSecret,
		Token:        *dbClient.Token,
		RefreshToken: *dbClient.RefreshToken,
		AthleteId:    *dbClient.AthletId,
	}, nil
}

func (s DatabaseServer) UpsertActivity(ctx context.Context, req *grpcStrava.Activity) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.Storer.UpsertActivity(req, false)
}

func (s DatabaseServer) GetActivities(ctx context.Context, req *grpcDB.ActivitiesRequest) (*grpcDB.ActivitiesResponse, error) {
	return s.Storer.GetActivities(req)
}

func (s DatabaseServer) UpsertActivityFromCSV(ctx context.Context, req *grpcStrava.Activity) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.Storer.UpsertActivity(req, true)
}
