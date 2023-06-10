package config

const Prefix = "/run"

type ServiceConfig struct {
	RunnerName         string `env:"RUNNER_NAME"`
	IsDev              bool   `env:"IS_DEV"`
	HttpGatewayPort    uint64 `env:"HTTP_GATEWAY_PORT"`
	RunningAppPort     uint64 `env:"RUNNING_APP_PORT"`
	RunningAppPassword string `env:"RUNNING_APP_PASSWORD"`
	JWTKey             string `env:"JWT_KEY"`
	MasterPassword     string `env:"MASTER_PASSWORD"`
}
