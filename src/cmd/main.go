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

//TODO: Надо поделить все на папки (handlers), но сейчас есть цикличные зависимости

const (
	envLocal = "local"
	envProd  = "prod"
)

// main is the entry point of the application. It loads the configuration, sets up the logger, initializes the database,
// sets up the API, starts the server, and handles graceful shutdown.
//
// No parameters.
// No return values.
func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting application", slog.String("env", cfg.Env))

	storage := postgres.MustInitDB(cfg.ConnectionString)

	log.Info("initialized db")

	setupAPI(storage, log, context.Background())

	log.Info("started server", slog.String("address", cfg.HTTPServer.Address))

	if err := http.ListenAndServe(cfg.HTTPServer.Address, nil); err != nil {
		log.Error("failed to start server", err)
	}

	//TODO: graceful shutdown
}

// setupAPI sets up the API endpoints for the application.
//
// Parameters:
// - storage: a pointer to a postgres.Postgres object representing the database connection.
// - log: a pointer to a slog.Logger object representing the logger.
// - ctx: a context.Context object representing the context.
//
// Return:
// None.
func setupAPI(storage *postgres.Postgres, log *slog.Logger, ctx context.Context) {
	m := ws.NewManager(storage, log, ctx)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	})

	http.HandleFunc("/ws", m.ServeWS)
	http.HandleFunc("/login", m.LoginHandler)
	http.HandleFunc("/register", m.RegisterUser)
}

// setupLogger initializes a logger based on the provided environment.
//
// Parameters:
// - env: a string representing the environment (e.g., "local", "prod").
//
// Returns:
// - *slog.Logger: a pointer to the initialized logger.
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
