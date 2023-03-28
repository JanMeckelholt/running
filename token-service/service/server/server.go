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
	return &emptypb.Empty{}, s.TokenStorer.UpsertClient(req.GetClientSecret(), req.GetToken(), req.GetRefreshToken(), req.GetClientId())
}

func (s TokenServer) GetClient(ctx context.Context, req *grpcToken.ClientId) (*grpcToken.Client, error) {
	athlete, err := s.TokenStorer.GetClient(req.GetClientId())
	return &grpcToken.Client{
		ClientId:     athlete.ClientId,
		ClientSecret: athlete.ClientSecret,
		Token:        athlete.Token,
		RefreshToken: athlete.RefreshToken,
	}, err
}
