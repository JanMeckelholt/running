package config

type ServiceConfig struct {
	StravaServiceName   string `env:"STRAVA_SERVICE_NAME"`
	DatabaseServiceName string `env:"DATABASE_SERVICE_NAME"`
	RunningAppPort      uint64 `env:"RUNNING_APP_PORT"`
}
