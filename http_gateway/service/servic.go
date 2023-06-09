package service

import (
	"time"

	"github.com/JanMeckelholt/running/http_gateway/service/clients"
	"github.com/JanMeckelholt/running/http_gateway/service/config"
)

const LoginRoute = config.Prefix + "/login"
const WebsiteRoute = config.Prefix + "/website"

type Service struct {
	Clients clients.Clients
	Config  config.ServiceConfig
}

type ClimbResponse struct {
	Climb uint64
}

type ActivitiesRequestBody struct {
	ClientId  *string
	Since     *uint64 `json:"since,omitempty"`
	Until     *uint64 `json:"until,omitempty"`
	WeekSince *int64  `json:"weekSince,omitempty"`
	WeekUntil *int64  `json:"weekUntil,omitempty"`
	Week      *int64  `json:"week,omitempty"`
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
	StartOfWeekUnix *uint64
	StartOfWeekStr  *string
	Distance        *uint64
	TimeUnix        *uint64
	TimeStr         *string
	NumberOfRuns    *uint64
	Climb           *uint64
	NumberOfOthers  *uint64
}

func (ws *WeekSummary) SetTimeStr() {
	t := time.Duration(*ws.TimeUnix * uint64(time.Second))
	timeStr := t.String()
	ws.TimeStr = &timeStr
}

type WeekSummariesResponse struct {
	WeekSummaries []*WeekSummary
}
