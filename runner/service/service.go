package service

import (
	"google.golang.org/grpc"
	"net/http"

	"running/runner/service/config"
)

type Service struct {
	GrpcServer    *grpc.Server
	HttpServer    *http.Server
	ServiceConfig config.ServiceConfig
}

type RunnerConfig struct {
}
