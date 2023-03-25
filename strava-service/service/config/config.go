package config

import "net/url"

type ServiceConfig struct {
	StravaToken string  `env:"STRAVA_TOKEN""`
	StravaURL   url.URL `env:"BASE_URL_STRAVA"`
}
