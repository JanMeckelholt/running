package server

import (
	"context"
	"net/http"
	"strings"

	"github.com/JanMeckelholt/running/common/grpc/database"
	"github.com/JanMeckelholt/running/common/grpc/runner"
	"github.com/JanMeckelholt/running/common/grpc/strava"
	"github.com/JanMeckelholt/running/common/models"
	"github.com/JanMeckelholt/running/http_gateway/service"
	"github.com/JanMeckelholt/running/http_gateway/service/auth"
	"github.com/JanMeckelholt/running/http_gateway/service/clients"
	regexphandler "github.com/JanMeckelholt/running/http_gateway/service/regexpHandler"
	"github.com/golang-jwt/jwt/v4"
)

type HttpGatewayServer struct {
	clients clients.Clients
}

func NewHttpGatewayServer(clients clients.Clients) (*HttpGatewayServer, error) {
	return &HttpGatewayServer{
		clients: clients,
	}, nil
}

func (rs HttpGatewayServer) CreateAthlete(ctx context.Context, request database.Athlete) error {
	_, err := rs.clients.RunnerClient.CreateAthlete(ctx, &request)
	return err
}

func (rs HttpGatewayServer) CreateClient(ctx context.Context, request database.Client) error {
	_, err := rs.clients.RunnerClient.CreateClient(ctx, &request)
	return err
}

func (rs HttpGatewayServer) GetAthlete(ctx context.Context, request runner.AthleteRequest) (*strava.AthleteResponse, error) {
	return rs.clients.RunnerClient.GetAthlete(ctx, &request)
}

func (rs HttpGatewayServer) ActivitiesToDB(ctx context.Context, request runner.ActivitiesRequest) error {
	_, err := rs.clients.RunnerClient.ActivitiesToDB(ctx, &request)
	return err
}

func (rs HttpGatewayServer) GetActivities(ctx context.Context, request database.ActivitiesRequest) (*database.ActivitiesResponse, error) {
	return rs.clients.RunnerClient.GetActivities(ctx, &request)
}

func (rs HttpGatewayServer) GetWeekSummaries(ctx context.Context, request runner.WeekSummariesRequest) (*runner.WeekSummariesResponse, error) {
	return rs.clients.RunnerClient.GetWeekSummaries(ctx, &request)
}

func (rs HttpGatewayServer) GetWeekSummary(ctx context.Context, request runner.WeekSummaryRequest) (*runner.WeekSummary, error) {
	return rs.clients.RunnerClient.GetWeekSummary(ctx, &request)
}

func AuthMiddleware(next regexphandler.RegexpHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// unprotected endpoints
		if r.URL.Path == service.LoginRoute {
			next.ServeHTTP(w, r)
			return
		}
		if strings.ToLower(r.URL.Path[:5]) == service.JungRoute {
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
