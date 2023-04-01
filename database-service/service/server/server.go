package server

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/JanMeckelholt/running/database-service/service"

	grpcDB "github.com/JanMeckelholt/running/common/grpc/database"
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
	return &emptypb.Empty{}, s.Storer.UpsertClient(req.GetClientId(), req.GetClientSecret(), req.GetToken(), req.GetRefreshToken())
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
	}, nil
}
