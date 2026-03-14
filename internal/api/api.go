package api

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

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
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	defer stop()

	server := &http.Server{
		Addr: ":" + a.cfg.Port,
		Handler: a.router,
	}

	go func() {

		log.Printf("server started on port %s", a.cfg.Port)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Panicf("error starting the server %v", err)
		}


	} ()

	<- ctx.Done()
	stop()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
	defer cancel()

	err := server.Shutdown(shutdownCtx)

	if err != nil {
		log.Printf("error shutting down the server gracefully: %v", err)
		time.Sleep(time.Second * 3)
	}

	log.Print("server shutdown gracefully")

	return err
}
