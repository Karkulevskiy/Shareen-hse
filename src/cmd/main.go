package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/karkulevskiy/shareen/src/internal/config"
	"github.com/karkulevskiy/shareen/src/internal/storage/postgres"
	"github.com/karkulevskiy/shareen/src/internal/ws"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting application", slog.String("env", cfg.Env))

	storage := postgres.MustInitDB(cfg.ConnectionString)

	log.Info("initialized db")

	setupAPI(storage, log)

	if err := http.ListenAndServe(cfg.HTTPServer.Address, nil); err != nil {
		log.Error("failed to start server", err)
	}

	//TODO: graceful shutdown
}

func setupAPI(storage *postgres.Storage, log *slog.Logger) {
	manager := ws.NewManager(storage, log)

	http.HandleFunc("/ws", manager.serveWS)
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
