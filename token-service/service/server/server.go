package server

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/JanMeckelholt/running/token-service/service"

	grpcToken "github.com/JanMeckelholt/running/common/grpc/token"
)

type TokenServer struct {
	TokenStorer *service.TokenStorer
	grpcToken.UnimplementedTokenServer
}

func NewServer(storer *service.TokenStorer) *TokenServer {

	ts := new(TokenServer)
	ts.TokenStorer = storer
	return ts
}

func (s TokenServer) UpsertClient(ctx context.Context, req *grpcToken.Client) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.TokenStorer.UpsertClient(req.GetClientId(), req.GetClientSecret(), req.GetToken(), req.GetRefreshToken())
}

func (s TokenServer) UpdateClient(ctx context.Context, req *grpcToken.UpdateRequest) (*grpcToken.Client, error) {
	return s.TokenStorer.UpdateClient(req.GetClientId(), req.GetKvPairs())
}

func (s TokenServer) GetClient(ctx context.Context, req *grpcToken.ClientId) (*grpcToken.Client, error) {
	dbClient, err := s.TokenStorer.GetDBClient(req.GetClientId())
	if err != nil {
		return nil, err
	}
	if dbClient == nil {
		return nil, nil
	}
	return &grpcToken.Client{
		ClientId:     *dbClient.ClientId,
		ClientSecret: *dbClient.ClientSecret,
		Token:        *dbClient.Token,
		RefreshToken: *dbClient.RefreshToken,
	}, nil
}
