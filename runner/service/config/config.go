package config

type ServiceConfig struct {
	StravaServiceName   string `env:"STRAVA_SERVICE_NAME"`
	DatabaseServiceName string `env:"TOKEN_SERVICE_NAME"`
}
