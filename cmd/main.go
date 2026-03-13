package main

import (
	"context"
	"log"

	"github.com/kratoio/krato/internal/config"
	"github.com/kratoio/krato/internal/database"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		log.Fatalf("error loading config: %s", err)
	}

	dbpool, err := database.Connect(*cfg)

	if err != nil {
		log.Fatalf("error connecting to database: %s", err)
	}

	defer dbpool.Close()

	if err := dbpool.Ping(context.Background()); err != nil {
		log.Fatalf("error running sql query: %s", err)
	}
}
