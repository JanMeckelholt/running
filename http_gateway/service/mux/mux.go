package mux

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/JanMeckelholt/running/common/grpc/database"
	"github.com/JanMeckelholt/running/common/grpc/runner"
	"github.com/JanMeckelholt/running/common/health"
	"github.com/JanMeckelholt/running/common/utils"
	"github.com/JanMeckelholt/running/http_gateway/service"
	"github.com/JanMeckelholt/running/http_gateway/service/auth"
	"github.com/JanMeckelholt/running/http_gateway/service/server"
)

func Handler(uri string, s *server.HttpGatewayServer) http.Handler {
	switch uri {
	case service.LoginRoute:
		return auth.LoginHandler()
	case service.WebsiteRoute:
		return auth.WebsiteHandler()
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
				res, err := s.GetRunner(context.Background(), runner.RunnerRequest{ClientId: rB.ClientId})
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
				log.Infof("AthletId: %d", uint64(rB.AthletId))
				err = s.CreateClient(context.Background(), database.Client{
					ClientId:     rB.ClientId,
					ClientSecret: rB.ClientSecret,
					Token:        rB.Token,
					RefreshToken: rB.RefreshToken,
					AthleteId:    uint64(rB.AthletId),
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
				if rB.WeekSince != nil && rB.YearSince != nil {
					log.Infof("using sinceWeek/siceYear-parameter, since-parameter is overwritten")
					s := utils.GetStartOfWeek(*rB.YearSince, *rB.WeekSince)
					rB.Since = &s
					if rB.WeekUntil == nil || rB.YearUntil == nil {
						rB.Until = nil
					} else {
						until := utils.GetStartOfWeek(*rB.YearUntil, *rB.WeekUntil) + 7*24*60*60
						rB.Until = &until
					}
				}
				log.Infof("Getting activities for client %s since %d until %d", *rB.ClientId, rB.Since, *rB.Until)
				res, err := s.GetActivities(context.Background(), database.ActivitiesRequest{Since: *rB.Since, Until: *rB.Until, ClientId: *rB.ClientId})
				//res, err := rs.GetActivities(context.Background(), strava.ActivityRequest{Token: rB.Token, Since: rB.Since, ClientId: rB.ClientId})
				if err != nil {
					rw.WriteHeader(http.StatusInternalServerError)
					_ = json.NewEncoder(rw).Encode(fmt.Sprintf("Error requesting StravaClient %s", err.Error()))
					return
				}
				rw.WriteHeader(http.StatusOK)
				_ = json.NewEncoder(rw).Encode(res.GetActivities())
			case http.MethodOptions:
				rw.Header().Set("Allow", "OPTIONS, POST")
				rw.Header().Set("Cache-Control", "max-age=604800")
				rw.WriteHeader(http.StatusOK)
			default:
				rw.WriteHeader(http.StatusMethodNotAllowed)
			}
		})
	case "/weeksummary":
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.Header().Add("Content-Type", "application/json")
			switch req.Method {
			case http.MethodGet:
				var weekInt, yearInt int
				weekStr := req.URL.Query().Get("week")
				yearStr := req.URL.Query().Get("year")

				weekInt, err := strconv.Atoi(weekStr)
				if err != nil {
					weekInt, yearInt = time.Now().ISOWeek()
				} else {
					yearInt, err = strconv.Atoi(yearStr)
					if err != nil {
						weekInt, yearInt = time.Now().ISOWeek()
					}
				}
				log.Info("get weeks")
				client := req.URL.Query().Get("client")
				weeksummarryResponse, err := s.GetWeekSummary(context.Background(), runner.WeekSummaryRequest{ClientId: client, Week: uint64(weekInt), Year: uint64(yearInt)})
				if err != nil {
					rw.WriteHeader(http.StatusInternalServerError)
					_ = json.NewEncoder(rw).Encode(fmt.Sprintf("Error requesting WeekSummary %s", err.Error()))
					return
				}
				rw.WriteHeader(http.StatusOK)
				_ = json.NewEncoder(rw).Encode(weeksummarryResponse)
			case http.MethodOptions:
				rw.Header().Set("Allow", "OPTIONS, GET")
				rw.Header().Set("Cache-Control", "max-age=604800")
				rw.WriteHeader(http.StatusOK)
			default:
				rw.WriteHeader(http.StatusMethodNotAllowed)
			}
		})
	case "/weeksummaries":
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.Header().Add("Content-Type", "application/json")
			switch req.Method {
			case http.MethodGet:
				log.Info("get weeksummaries")
				client := req.URL.Query().Get("client")
				weekSinceStr := req.URL.Query().Get("weekSince")
				weekUntilStr := req.URL.Query().Get("weekUntil")
				weekSince, err := strconv.ParseUint(weekSinceStr, 10, 64)
				if err != nil {
					weekSince = 0
				}
				weekUntil, err := strconv.ParseUint(weekUntilStr, 10, 64)
				if err != nil {
					weekUntil = 0
				}
				weeksummariesResponse, err := s.GetWeekSummaries(context.Background(), runner.WeekSummariesRequest{ClientId: client, WeekSince: weekSince, WeekUntil: weekUntil})
				if err != nil {
					rw.WriteHeader(http.StatusInternalServerError)
					_ = json.NewEncoder(rw).Encode(fmt.Sprintf("Error requesting WeekSummaries %s", err.Error()))
					return
				}
				rw.WriteHeader(http.StatusOK)
				_ = json.NewEncoder(rw).Encode(weeksummariesResponse)

			case http.MethodPost:
				log.Info("get weeksummaries")
				rB := service.ActivitiesRequestBody{}
				err := json.NewDecoder(req.Body).Decode(&rB)
				if err != nil {
					log.Errorf(err.Error())
				}
				if rB.WeekSince != nil && rB.YearSince != nil {
					log.Infof("using sinceWeek/siceYear-parameter, since-parameter is overwritten")
					s := utils.GetStartOfWeek(*rB.YearSince, *rB.WeekSince)
					rB.Since = &s
					if rB.WeekUntil == nil || rB.YearUntil == nil {
						rB.Until = nil
					} else {
						until := utils.GetStartOfWeek(*rB.YearUntil, *rB.WeekUntil) + 7*24*60*60
						rB.Until = &until
					}
				}
				weeksummariesResponse, err := s.GetWeekSummaries(context.Background(), runner.WeekSummariesRequest{ClientId: *rB.ClientId, WeekSince: *rB.WeekSince, WeekUntil: *rB.WeekUntil})
				if err != nil {
					rw.WriteHeader(http.StatusInternalServerError)
					_ = json.NewEncoder(rw).Encode(fmt.Sprintf("Error requesting WeekSummaries %s", err.Error()))
					return
				}
				rw.WriteHeader(http.StatusOK)
				_ = json.NewEncoder(rw).Encode(weeksummariesResponse)
			case http.MethodOptions:
				rw.Header().Set("Allow", "OPTIONS, GET, POST")
				rw.Header().Set("Cache-Control", "max-age=604800")
				rw.WriteHeader(http.StatusOK)
			default:
				rw.WriteHeader(http.StatusMethodNotAllowed)
			}
		})
	case "/activitiesToDB":
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.Header().Add("Content-Type", "application/json")
			switch req.Method {
			case http.MethodPost:
				log.Infof("Getting activites from Strava for last weeks and saving in DB")
				rB := service.ActivitiesRequestBody{}
				err := json.NewDecoder(req.Body).Decode(&rB)
				if err != nil {
					log.Errorf(err.Error())
				}
				if rB.WeekSince != nil && rB.YearSince != nil {
					log.Infof("using sinceWeek/siceYear-parameter, since-parameter is overwritten")
					s := utils.GetStartOfWeek(*rB.YearSince, *rB.WeekSince)
					rB.Since = &s
				}
				err = s.ActivitiesToDB(context.Background(), runner.ActivitiesRequest{Since: *rB.Since, ClientId: *rB.ClientId})
				if err != nil {
					rw.WriteHeader(http.StatusInternalServerError)
					_ = json.NewEncoder(rw).Encode(fmt.Sprintf("Error adding Activites to DB %s", err.Error()))
					return
				}
				rw.WriteHeader(http.StatusOK)
				_ = json.NewEncoder(rw).Encode("Added Activities to DB")
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
