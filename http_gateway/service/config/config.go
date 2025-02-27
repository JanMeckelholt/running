package config

const RunPrefix = ""
const ApiPrefix = "/api"

type ServiceConfig struct {
	RunnerName         string `env:"RUNNER_NAME"`
	IsDev              bool   `env:"IS_DEV"`
	RunningAppPort     uint64 `env:"RUNNING_APP_PORT"`
	RunningAppPassword string `env:"RUNNING_APP_PASSWORD"`
	JWTKey             string `env:"JWT_KEY"`
	MasterPassword     string `env:"MASTER_PASSWORD"`
}
