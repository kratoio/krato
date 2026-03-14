package main

import (
	"os"

	"github.com/kratoio/krato/internal/api"
	"github.com/kratoio/krato/internal/config"
	"github.com/kratoio/krato/internal/logger"
)

func main() {
	l := logger.New("dev")

	cfg, err := config.Load()

	if err != nil {
		l.Error("error loading env file", "err", err)
		os.Exit(1)
	}

	l = logger.New(cfg.Env)

	a, err := api.New(l, cfg)

	if err != nil {
		l.Error("error creating api instance", "err", err)
		os.Exit(1)
	}

	if err := a.Start(); err != nil {
		l.Error("error starting the server", "err", err)
		os.Exit(1)
	}
}
