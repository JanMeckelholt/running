package config

type ServiceConfig struct {
	StravaURL        string `env:"BASE_URL_STRAVA"`
	TokenServiceName string `env:"TOKEN_SERVICE_NAME"`
}
