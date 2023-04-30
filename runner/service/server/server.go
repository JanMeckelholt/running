package server

import (
	"context"
	"time"

	"github.com/JanMeckelholt/running/common/grpc/database"
	"github.com/JanMeckelholt/running/common/grpc/runner"
	"github.com/JanMeckelholt/running/common/grpc/strava"
	"github.com/JanMeckelholt/running/common/utils"
	"github.com/JanMeckelholt/running/runner/service/clients"
	"github.com/JanMeckelholt/running/runner/service/logic"
	"google.golang.org/protobuf/types/known/emptypb"

	log "github.com/sirupsen/logrus"
)

type RunnerServer struct {
	clients clients.Clients
	runner.UnimplementedRunnerServer
}

func NewRunnerServer(clients clients.Clients) (*RunnerServer, error) {
	return &RunnerServer{
		clients: clients,
	}, nil
}

func (rs RunnerServer) GetRunner(ctx context.Context, request *runner.RunnerRequest) (*strava.RunnerResponse, error) {
	return rs.clients.StravaClient.GetRunner(ctx, &strava.RunnerRequest{ClientId: request.GetClientId()})
}

func (rs RunnerServer) GetActivitiesFromStrava(ctx context.Context, request strava.ActivitiesRequest) (*strava.ActivitiesResponse, error) {
	return rs.clients.StravaClient.GetActivities(ctx, &request)
}

func (rs RunnerServer) ActivitiesToDB(ctx context.Context, request *runner.ActivitiesRequest) (*emptypb.Empty, error) {
	clientId := request.GetClientId()
	dbClient, err := rs.clients.DatabaseClient.GetClient(ctx, &database.ClientId{ClientId: clientId})
	if err != nil {
		return nil, err
	}
	activitiesResp, err := rs.clients.StravaClient.GetActivities(ctx, &strava.ActivitiesRequest{ClientId: clientId, Token: dbClient.Token, Since: request.GetSince()})
	if err != nil {
		return nil, err
	}
	log.Infof("got %d activities from Strava", len(activitiesResp.GetActivities()))
	for i, activity := range activitiesResp.GetActivities() {
		log.Infof("activity %d: %s", i, activity.GetName())
		_, err := rs.clients.DatabaseClient.UpsertActivity(ctx, activity)
		if err != nil {
			return nil, err
		}
	}
	return &emptypb.Empty{}, nil
}

func (rs RunnerServer) CreateClient(ctx context.Context, request *database.Client) (*emptypb.Empty, error) {
	return rs.clients.DatabaseClient.UpsertClient(ctx, request)
}

func (rs RunnerServer) GetActivities(ctx context.Context, request *database.ActivitiesRequest) (*database.ActivitiesResponse, error) {
	return rs.clients.DatabaseClient.GetActivities(ctx, request)
}

func (rs RunnerServer) GetWeekSummaries(ctx context.Context, request *runner.WeekSummariesRequest) (*runner.WeekSummariesResponse, error) {
	startOfFirstWeek := utils.GetStartOfWeek(request.GetWeekSince())
	var endOfLastWeek uint64
	if request.GetWeekUntil() == 0 {
		endOfLastWeek = uint64(time.Now().Unix())
	} else {
		endOfLastWeek = startOfFirstWeek + 7*24*60*60*uint64((1-(request.GetWeekSince()-request.GetWeekUntil())))
	}
	res, err := rs.clients.DatabaseClient.GetActivities(context.Background(), &database.ActivitiesRequest{Since: startOfFirstWeek, Until: endOfLastWeek, ClientId: request.GetClientId()})
	if err != nil {
		return nil, err
	}
	weeksummarryResponse := logic.GetWeekSummariesResponse(res, startOfFirstWeek, endOfLastWeek)

	return &weeksummarryResponse, nil
}

func (rs RunnerServer) GetWeekSummary(ctx context.Context, request *runner.WeekSummaryRequest) (*runner.WeekSummary, error) {
	startOfWeek := utils.GetStartOfWeek(request.GetWeek())
	endOfWeek := startOfWeek + 7*24*60*60
	res, err := rs.clients.DatabaseClient.GetActivities(context.Background(), &database.ActivitiesRequest{Since: startOfWeek, Until: endOfWeek, ClientId: request.GetClientId()})
	if err != nil {
		return nil, err
	}
	week := logic.GetWeek(res, startOfWeek)

	return &week, nil
}

func (rs RunnerServer) StravaActivitiesToDB(ctx context.Context, request strava.ActivitiesRequest) error {
	_, err := rs.clients.StravaClient.ActivitiesToDB(ctx, &request)
	return err
}

func (rs RunnerServer) Health(ctx context.Context, _ *runner.HealthMessage) (*runner.HealthMessage, error) {
	log.Info("Calling Health :-)")
	return &runner.HealthMessage{Health: "OK"}, nil
}
