package config

type ServiceConfig struct {
	RunnerName     string `env:"RUNNER_NAME"`
	IsDev          bool   `env:"IS_DEV"`
	RunningAppPort uint64 `env:"RUNNING_APP_PORT"`
}
