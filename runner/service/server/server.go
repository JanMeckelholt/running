package server

import (
	"context"

	"github.com/JanMeckelholt/running/common/grpc/strava"
	"github.com/JanMeckelholt/running/common/grpc/token"
	"github.com/JanMeckelholt/running/runner/service/clients"
)

type RunnerServer struct {
	clients clients.Clients
}

func NewRunnerServer(clients clients.Clients) (*RunnerServer, error) {
	return &RunnerServer{
		clients: clients,
	}, nil
}

func (rs RunnerServer) GetAthlet(ctx context.Context, request strava.RunnerRequest) (*strava.RunnerResponse, error) {
	return rs.clients.StravaClient.GetRunner(ctx, &request)
}

func (rs RunnerServer) GetActivities(ctx context.Context, request strava.ActivityRequest) (*strava.ActivitiesResponse, error) {
	return rs.clients.StravaClient.GetActivities(ctx, &request)
}

func (rs RunnerServer) CreateRunner(ctx context.Context, request token.Client) error {
	_, err := rs.clients.TokenClient.UpsertClient(ctx, &request)
	return err
}
