package service

import (
	"github.com/JanMeckelholt/running/populate_db/service/clients"
	"github.com/JanMeckelholt/running/populate_db/service/config"
)

type Service struct {
	ServiceConfig config.ServiceConfig
	Clients       clients.Clients
}
