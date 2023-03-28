package config

type StorerConfig struct {
	PostgresUser     string `env:"POSTGRES_USER"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresDB       string `env:"POSTGRES_DB"`
	PORT             int    `env:"POSTGRES_PORT"`
	HOST             string `env:"POSTGRES_HOST"`
}
