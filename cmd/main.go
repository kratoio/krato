package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
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

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	if err = http.ListenAndServe(":" + cfg.Port, r); err != nil {
		log.Fatalf("error listening to the port %s: %s", cfg.Port, err)
	}

}
