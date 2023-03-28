package service

import (
	"github.com/JanMeckelholt/running/strava-service/service/clients"
	"github.com/JanMeckelholt/running/strava-service/service/config"
)

type Service struct {
	ServiceConfig config.ServiceConfig
	Clients       clients.Clients
}
