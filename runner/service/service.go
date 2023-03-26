package service

import (
	"github.com/JanMeckelholt/running/runner/service/clients"
)

type Service struct {
	Clients clients.Clients
}

type RunnerConfig struct {
}

type RunnerRequestBody struct {
	Token string
}
