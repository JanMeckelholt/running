package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/JanMeckelholt/running/common/grpc/database"
	"github.com/JanMeckelholt/running/common/grpc/runner"
	"github.com/JanMeckelholt/running/common/grpc/strava"
	"github.com/JanMeckelholt/running/common/models"
	"github.com/JanMeckelholt/running/http_gateway/service"
	"github.com/JanMeckelholt/running/http_gateway/service/auth"
	"github.com/JanMeckelholt/running/http_gateway/service/clients"
	"github.com/JanMeckelholt/running/http_gateway/service/config"
	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"
)

type HttpGatewayServer struct {
	clients clients.Clients
}

func NewHttpGatewayServer(clients clients.Clients) (*HttpGatewayServer, error) {
	return &HttpGatewayServer{
		clients: clients,
	}, nil
}

func (rs HttpGatewayServer) GetRunner(ctx context.Context, request runner.RunnerRequest) (*strava.RunnerResponse, error) {
	return rs.clients.RunnerClient.GetRunner(ctx, &request)
}

func (rs HttpGatewayServer) ActivitiesToDB(ctx context.Context, request runner.ActivitiesRequest) error {
	_, err := rs.clients.RunnerClient.ActivitiesToDB(ctx, &request)
	return err
}

func (rs HttpGatewayServer) CreateClient(ctx context.Context, request database.Client) error {
	_, err := rs.clients.RunnerClient.CreateClient(ctx, &request)
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

func CorsMiddleware(next http.Handler, config config.ServiceConfig) http.Handler {
	allowList := map[string]bool{
		"janmeckelholt.de:444":                                           true,
		"https://janmeckelholt.de:444":                                   true,
		"https://janmeckelholt.de:444/run":                               true,
		"https://janmeckelholt.de:444/run/":                              true,
		"http://janmeckelholt.de":                                        true,
		"http://janmeckelholt.de/run":                                    true,
		"http://janmeckelholt.de/run/":                                   true,
		fmt.Sprintf("http://homepage-server:%d", config.RunningAppPort):  true,
		fmt.Sprintf("https://homepage-server:%d", config.RunningAppPort): true,
		fmt.Sprintf("http://localhost:%d", config.RunningAppPort):        true,
		fmt.Sprintf("https://localhost:%d", config.RunningAppPort):       true,
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); allowList[origin] {
			log.Infof("allowing Origin: %s", origin)
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			log.Infof("not allowing Origin: %s", r.Header.Get(("Origin")))
		}

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
