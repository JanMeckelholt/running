package controller

import (
	"context"

	"github.com/JanMeckelholt/running/common/grpc/strava"
)

type RunnerController struct {
	client strava.StravaClient
}

func NewRunnerController(client strava.StravaClient) RunnerController {
	return RunnerController{
		client: client,
	}
}

func (rc RunnerController) GetAthlet(ctx context.Context) (string, error) {
	req := strava.Request{
		Name: "Jan",
	}
	resp, err := rc.client.Get(ctx, &req)
	if err != nil {
		return "", err
	}
	return resp.Something, nil

}
