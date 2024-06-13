package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/karkulevskiy/shareen/src/internal/config"
	"github.com/karkulevskiy/shareen/src/internal/storage/postgres"
	"github.com/karkulevskiy/shareen/src/internal/ws"
)

//TODO: CORS policy

const (
	envLocal = "local"
	envProd  = "prod"
)

// main starts application
func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)

	log.Info("starting application", slog.String("env", cfg.Env))

	storage := postgres.MustInit(cfg.ConnectionString)

	log.Info("initialized Database")

	setupAPI(storage, log, context.Background())

	log.Info("started server", slog.String("address", cfg.HTTPServer.Address))

	if err := http.ListenAndServe(cfg.HTTPServer.Address, nil); err != nil {
		log.Error("failed to start server", err)
	}

	//TODO: graceful shutdown
}

// setupAPI sets up http handlers
func setupAPI(storage *postgres.Postgres, log *slog.Logger, ctx context.Context) {
	m := ws.NewManager(storage, log, ctx)

	http.HandleFunc("/ws", m.ServeWS)
	http.HandleFunc("/login", m.LoginHandler)
	http.HandleFunc("/register", m.RegisterUser)
}

// setupLogger sets up logger
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
