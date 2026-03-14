package logger

import (
	"log/slog"
	"os"
)

func New(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case "dev":
		log = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug, AddSource: true}))
	default:
		log = slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo, AddSource: false}))
	}

	return log
}
