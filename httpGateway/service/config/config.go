package config

type ServiceConfig struct {
	RunnerGrpcName string `env:"RUNNERGRPC_NAME"`
	RunnerName     string `env:"RUNNER_NAME"`
	RunningAppPort uint64 `env:"RUNNING_APP_PORT"`
}
