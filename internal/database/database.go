package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kratoio/krato/internal/config"
)

func Connect(cfg config.Config) (*pgxpool.Pool, error) {
	DatabaseURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode)

	pool, err := pgxpool.New(context.Background(), DatabaseURL)

	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %s", err)
	}

	return pool, nil
}
