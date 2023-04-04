package config

type ServiceConfig struct {
	Enabled             bool   `env:"ENABLED"`
	DatabaseServiceName string `env:"DATABASE_SERVICE_NAME"`
	ClientsConfig       clientConfig
}

type clientConfig struct {
	ClientSecret string `env:"CLIENT_SECRET"`
	Token        string `env:"CLIENT_TOKEN"`
	RefreshToken string `env:"CLIENT_REFRESHTOKEN"`
	AthletId     uint64 `env:"CLIENT_ATHLETEID"`
	ClientId     string `env:"CLIENT_ID"`
}
