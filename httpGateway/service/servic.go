package service

import (
	"time"

	"github.com/JanMeckelholt/running/httpGateway/service/clients"
	"github.com/JanMeckelholt/running/httpGateway/service/config"
)

const LoginRoute = "/login"

type Service struct {
	Clients clients.Clients
	Config  config.ServiceConfig
}

type ClimbResponse struct {
	Climb uint64
}

type ActivitiesRequestBody struct {
	ClientId string
	Since    uint64 `json:"since,omitempty"`
}

type RunnerRequestBody struct {
	ClientId string
}

type RunnerCreateBody struct {
	Token        string
	RefreshToken string
	ClientId     string
	ClientSecret string
	AthletId     uint64
}

type WeekSummary struct {
	Distance       *uint64
	TimeUnix       *uint64
	TimeStr        *string
	NumberOfRuns   *uint64
	Climb          *uint64
	NumberOfOthers *uint64
}

func (ws *WeekSummary) SetTimeStr() {
	t := time.Duration(*ws.TimeUnix * uint64(time.Second))
	timeStr := t.String()
	ws.TimeStr = &timeStr
}

type WeekSummaryResponse struct {
	WeekSummaries []*WeekSummary
}
