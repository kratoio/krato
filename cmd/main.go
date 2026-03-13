package main

import (
	"log"

	"github.com/kratoio/krato/internal/api"
	"github.com/kratoio/krato/internal/config"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		log.Fatalf("error loading config: %s", err)
	}

	a, err := api.New(cfg)

	if err != nil {
		log.Fatalf("error creating api instance: %s", err)
	}

	if err := a.Start(); err != nil {
		log.Fatalf("error listening to the port %s: %s", cfg.Port, err)
	}
}
