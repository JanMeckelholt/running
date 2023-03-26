package mux

import (
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/JanMeckelholt/running/common/grpc/strava"
	"github.com/JanMeckelholt/running/common/health"
	"github.com/JanMeckelholt/running/runner/service/server"
)

func Handler(uri string, rs *server.RunnerServer) http.Handler {
	switch uri {
	case "/health":
		return health.Handler(health.Health{ServiceName: "runner"})

	case "/athlet":
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.Header().Add("Content-Type", "application/json")
			switch req.Method {
			case http.MethodPost:
				type RequestBody struct {
					Token string
				}
				rB := RequestBody{}
				err := json.NewDecoder(req.Body).Decode(&rB)
				if err != nil {
					log.Errorf(err.Error())
				}
				res, err := rs.GetAthlet(context.Background(), strava.RunnerRequest{Token: rB.Token})
				if err != nil {
					rw.WriteHeader(http.StatusInternalServerError)
					_ = json.NewEncoder(rw).Encode(fmt.Sprintf("Error requesting Strava %s", err.Error()))
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
