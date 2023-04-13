package server

import (
	"context"
	"net/http"

	"github.com/JanMeckelholt/running/common/grpc/database"
	"github.com/JanMeckelholt/running/common/grpc/runner"
	"github.com/JanMeckelholt/running/common/grpc/strava"
	"github.com/JanMeckelholt/running/httpGateway/service/clients"
)

type HTTPGatewayServer struct {
	clients clients.Clients
}

func NewHTTPGatewayServer(clients clients.Clients) (*HTTPGatewayServer, error) {
	return &HTTPGatewayServer{
		clients: clients,
	}, nil
}

func (rs HTTPGatewayServer) GetRunner(ctx context.Context, request strava.RunnerRequest) (*strava.RunnerResponse, error) {
	return rs.clients.RunnerClient.GetRunner(ctx, &request)
}

func (rs HTTPGatewayServer) ActivitiesToDB(ctx context.Context, request runner.ActivitiesRequest) error {
	_, err := rs.clients.RunnerClient.ActivitiesToDB(ctx, &request)
	return err
}

func (rs HTTPGatewayServer) CreateClient(ctx context.Context, request database.Client) error {
	_, err := rs.clients.RunnerClient.CreateClient(ctx, &request)
	return err
}

func (rs HTTPGatewayServer) GetActivities(ctx context.Context, request database.ActivitiesRequest) (*database.ActivitiesResponse, error) {
	return rs.clients.RunnerClient.GetActivities(ctx, &request)
}

func (rs HTTPGatewayServer) GetWeekSummaries(ctx context.Context, request runner.WeekSummariesRequest) (*runner.WeekSummariesResponse, error) {
	return rs.clients.RunnerClient.GetWeekSummaries(ctx, &request)
}

func CorsMiddleware(next http.Handler, allowOrigin string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		next.ServeHTTP(w, r)
	})
}
