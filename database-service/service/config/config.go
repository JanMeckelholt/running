package config

type StorerConfig struct {
	PostgresUser     string `env:"POSTGRES_USER"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresDB       string `env:"POSTGRES_DB"`
	PostgresPort     int    `env:"POSTGRES_PORT"`
	PostgresHost     string `env:"POSTGRES_HOST"`
	PostgresSSLMode  string `env:"POSTGRES_SSLMODE"`
}
