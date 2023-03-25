package config

type ServiceConfig struct {
	APIEndpoint string `env:"API_ENDPOINT" envDefault:"https://runnerserver/..."`
}

func DefaultServiceConfig() ServiceConfig {
	return ServiceConfig{
		APIEndpoint: "https://runnerserver/...",
	}
}
