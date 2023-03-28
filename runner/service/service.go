package service

import (
	"github.com/JanMeckelholt/running/runner/service/clients"
	"github.com/JanMeckelholt/running/runner/service/config"
)

type Service struct {
	Clients clients.Clients
	Config  config.ServiceConfig
}

type RunnerConfig struct {
}

type ClimbResponse struct {
	Climb int64
}

type RunnerRequestBody struct {
	Token        string
	RefreshToken string `json:"RefreshToken,omitempty"`
	Since        int64  `json:"since,omitempty"`
}

type RefreshTokenResponse struct {
	Refresh_token string
}

type RunnerCreateBody struct {
	Token        string
	RefreshToken string
	ClientId     string
	ClientSecret string
}
