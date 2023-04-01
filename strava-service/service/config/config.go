package config

type ServiceConfig struct {
	StravaURL           string `env:"BASE_URL_STRAVA"`
	DatabaseServiceName string `env:"DATABASE_SERVICE_NAME"`
}
