package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kratoio/krato/internal/config"
	"github.com/kratoio/krato/internal/database"
)

type API struct {
	cfg *config.Config
	router chi.Router
	dbpool *pgxpool.Pool
}

func New(cfg *config.Config) (*API, error) {
	dbpool, err := database.Connect(*cfg)

	if err != nil {
		return nil, err
	}

	r := chi.NewRouter()

	api := &API{
		cfg: cfg,
		dbpool: dbpool,
		router: r,
	}

	api.setupRoutes()

	return api, nil
}

func(a *API) Start() error {
	if err := http.ListenAndServe(":" + a.cfg.Port, a.router); err != nil {
		return err
	}
	return nil
}
