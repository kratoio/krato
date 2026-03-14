package api

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kratoio/krato/internal/config"
	"github.com/kratoio/krato/internal/database"
)

type API struct {
	l      *slog.Logger
	cfg    *config.Config
	router chi.Router
	dbpool *pgxpool.Pool
}

func New(l *slog.Logger, cfg *config.Config) (*API, error) {
	dbpool, err := database.Connect(*cfg)

	if err != nil {
		return nil, err
	}

	r := chi.NewRouter()

	api := &API{
		l:      l,
		cfg:    cfg,
		dbpool: dbpool,
		router: r,
	}

	api.setupRoutes()

	return api, nil
}

func (a *API) Start() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	defer stop()

	server := &http.Server{
		Addr:    ":" + a.cfg.Port,
		Handler: a.router,
	}

	go func() {

		a.l.Info("server started", "port", a.cfg.Port)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.l.Error("error starting the server", "err", err)
			os.Exit(1)
		}

	}()

	<-ctx.Done()
	stop()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	err := server.Shutdown(shutdownCtx)

	if err != nil {
		a.l.Error("error shutting down the server", "err", err)
		time.Sleep(time.Second * 3)
	}

	a.l.Info("server shutdown successfully")

	return err
}
