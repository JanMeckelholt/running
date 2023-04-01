package mux

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/JanMeckelholt/running/common/grpc/database"
	"github.com/JanMeckelholt/running/common/grpc/strava"
	"github.com/JanMeckelholt/running/common/health"
	"github.com/JanMeckelholt/running/runner/service"
	"github.com/JanMeckelholt/running/runner/service/logic"
	"github.com/JanMeckelholt/running/runner/service/server"
)

func Handler(uri string, rs *server.RunnerServer) http.Handler {
	switch uri {
	case "/health":
		return health.Handler(health.Health{ServiceName: "runner"})

	case "/athlete":
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.Header().Add("Content-Type", "application/json")
			switch req.Method {
			case http.MethodPost:
				rB := service.RunnerRequestBody{}
				err := json.NewDecoder(req.Body).Decode(&rB)
				if err != nil {
					log.Errorf(err.Error())
				}
				res, err := rs.GetAthlet(context.Background(), strava.RunnerRequest{Token: rB.Token})
				if err != nil {
					rw.WriteHeader(http.StatusInternalServerError)
					_ = json.NewEncoder(rw).Encode(fmt.Sprintf("Error requesting StravaClient %s", err.Error()))
					return
				}
				rw.WriteHeader(http.StatusOK)
				_ = json.NewEncoder(rw).Encode(res)
			case http.MethodOptions:
				rw.Header().Set("Allow", "OPTIONS, POST")
				rw.Header().Set("Cache-Control", "max-age=604800")
				rw.WriteHeader(http.StatusOK)
			default:
				rw.WriteHeader(http.StatusMethodNotAllowed)
			}
		})
	case "/athlete/create":
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.Header().Add("Content-Type", "application/json")
			switch req.Method {
			case http.MethodPost:
				rB := service.RunnerCreateBody{}
				err := json.NewDecoder(req.Body).Decode(&rB)
				if err != nil {
					log.Errorf(err.Error())
				}
				err = rs.CreateRunner(context.Background(), database.Client{
					ClientId:     rB.ClientId,
					ClientSecret: rB.ClientSecret,
					Token:        rB.Token,
					RefreshToken: rB.RefreshToken,
				})
				if err != nil {
					rw.WriteHeader(http.StatusInternalServerError)
					_ = json.NewEncoder(rw).Encode(fmt.Sprintf("Error creating runner %s", err.Error()))
					return
				}
				rw.WriteHeader(http.StatusOK)
				_ = json.NewEncoder(rw).Encode(fmt.Sprintf("runner successfully created!"))
			case http.MethodOptions:
				rw.Header().Set("Allow", "OPTIONS, POST")
				rw.Header().Set("Cache-Control", "max-age=604800")
				rw.WriteHeader(http.StatusOK)
			default:
				rw.WriteHeader(http.StatusMethodNotAllowed)
			}
		})
	case "/activities":
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.Header().Add("Content-Type", "application/json")
			switch req.Method {
			case http.MethodPost:
				rB := service.ActivitiesRequestBody{}
				err := json.NewDecoder(req.Body).Decode(&rB)
				if err != nil {
					log.Errorf(err.Error())
				}
				res, err := rs.GetActivities(context.Background(), strava.ActivityRequest{Token: rB.Token, Since: rB.Since})
				if err != nil {
					rw.WriteHeader(http.StatusInternalServerError)
					_ = json.NewEncoder(rw).Encode(fmt.Sprintf("Error requesting StravaClient %s", err.Error()))
					return
				}
				rw.WriteHeader(http.StatusOK)
				_ = json.NewEncoder(rw).Encode(res)
			case http.MethodOptions:
				rw.Header().Set("Allow", "OPTIONS, POST")
				rw.Header().Set("Cache-Control", "max-age=604800")
				rw.WriteHeader(http.StatusOK)
			default:
				rw.WriteHeader(http.StatusMethodNotAllowed)
			}
		})
	case "/weeklyclimb":
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.Header().Add("Content-Type", "application/json")
			switch req.Method {
			case http.MethodPost:
				rB := service.ActivitiesRequestBody{}
				err := json.NewDecoder(req.Body).Decode(&rB)
				if err != nil {
					log.Errorf(err.Error())
				}
				startoOfWeek := logic.GetStartOfWeek()
				res, err := rs.GetActivities(context.Background(), strava.ActivityRequest{Token: rB.Token, Since: startoOfWeek, ClientId: rB.ClientId})
				climb := logic.GetClimb(res)
				if err != nil {
					rw.WriteHeader(http.StatusInternalServerError)
					_ = json.NewEncoder(rw).Encode(fmt.Sprintf("Error requesting StravaClient %s", err.Error()))
					return
				}
				rw.WriteHeader(http.StatusOK)
				_ = json.NewEncoder(rw).Encode(service.ClimbResponse{Climb: climb})
			case http.MethodOptions:
				rw.Header().Set("Allow", "OPTIONS, POST")
				rw.Header().Set("Cache-Control", "max-age=604800")
				rw.WriteHeader(http.StatusOK)
			default:
				rw.WriteHeader(http.StatusMethodNotAllowed)
			}
		})
	}
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		switch req.Method {
		case http.MethodGet:
			rw.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(rw).Encode("not defined")
		case http.MethodOptions:
			rw.Header().Set("Allow", "OPTIONS, GET")
			rw.Header().Set("Cache-Control", "max-age=604800")
			rw.WriteHeader(http.StatusOK)
		default:
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}
