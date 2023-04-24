package server

import (
	"context"
	"net/http"

	"github.com/JanMeckelholt/running/common/grpc/database"
	"github.com/JanMeckelholt/running/common/grpc/runner"
	"github.com/JanMeckelholt/running/common/grpc/strava"
	"github.com/JanMeckelholt/running/common/models"
	"github.com/JanMeckelholt/running/httpGateway/service"
	"github.com/JanMeckelholt/running/httpGateway/service/auth"
	"github.com/JanMeckelholt/running/httpGateway/service/clients"
	"github.com/golang-jwt/jwt/v4"
)

type HTTPGatewayServer struct {
	clients clients.Clients
}

func NewHTTPGatewayServer(clients clients.Clients) (*HTTPGatewayServer, error) {
	return &HTTPGatewayServer{
		clients: clients,
	}, nil
}

func (rs HTTPGatewayServer) GetRunner(ctx context.Context, request runner.RunnerRequest) (*strava.RunnerResponse, error) {
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
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == service.LoginRoute || r.URL.Path == service.WebsiteRoute {
			next.ServeHTTP(w, r)
			return
		}
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		tknStr := c.Value
		claims := &models.Claims{}

		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return auth.JwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}