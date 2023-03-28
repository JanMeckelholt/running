package config

type ServiceConfig struct {
	StravaServiceName string `env:"STRAVA_SERVICE_NAME"`
	TokenServiceName  string `env:"TOKEN_SERVICE_NAME"`
}
