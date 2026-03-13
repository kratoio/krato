package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kratoio/krato/pkg/env"
)

type Config struct {
	Port string

	JWTSecret string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading env file: %s", err)
	}

	cfg := &Config{
		Port: env.Get("PORT", "8080"),
		JWTSecret: env.Get("JWT_SECRET", "secret"),
		DBHost: env.Get("DB_HOST", "localhost"),
		DBPort: env.Get("DB_PORT", "5432"),
		DBUser: env.Get("DB_USER", "postgres"),
		DBPassword: env.Get("DB_PASSWORD", ""),
		DBName: env.Get("DB_NAME", "db"),
		DBSSLMode: env.Get("DB_SSL_MODE", "disable"),
	}

	return cfg, nil
}
