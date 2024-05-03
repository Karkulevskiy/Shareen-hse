package main

import (
	"log/slog"
	"net/http"
	"os"

	_ "github.com/shareen/src/docs"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/shareen/src/internal/config"
	"github.com/shareen/src/internal/http-server/handlers"

	"github.com/shareen/src/internal/storage/postgres"
	_ "github.com/swaggo/files"                  // swagger embed files
	_ "github.com/swaggo/gin-swagger"            // gin-swagger middleware
	httpSwagger "github.com/swaggo/http-swagger" // http-swagger middleware
)

// log level
const (
	envLocal = "local"
	envProd  = "prod"
)

// @title           Shareen API
// @version         1.0
// @description     This is a sample server celler server.
// @license.name  Apache 2.0
// @host      localhost:8080
// @BasePath /
// @securityDefinitions.basic  BasicAuth
func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting application", slog.String("env", cfg.Env))

	storage := postgres.MustInitDB(cfg.ConnectionString)

	log.Info("initialized db")

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	//TODO: сделать проверку на валидность сайтов
	//TODO: доделать роуты, проверить, сделать для сайтов, сваггер, эндпоинты для обновления инфы в лобби
	router.Route("/lobby", func(r chi.Router) {
		r.Get("/{url}", handlers.GetLobby(log, storage.LobbyRepository))
		r.Post("/", handlers.CreateLobby(log, storage.LobbyRepository))
		r.Delete("/{url}", handlers.DeleteLobby(log, storage.LobbyRepository))
		r.Patch("/{url}-{video}", handlers.UpdateLobbyVideoURL(log, storage.LobbyRepository))
	})

	//TODO: как лучше передовать параметры? Строчкой или json???
	router.Route("/user", func(r chi.Router) {
		r.Get("/{id}", handlers.GetUser(log, storage.UserRepository))
		r.Post("/{name}", handlers.CreateUser(log, storage.UserRepository))
		r.Delete("/{id}", handlers.DeleteUser(log, storage.UserRepository))
		r.Patch("/{url}-{id}", handlers.JoinUserToLobby(log, storage.UserRepository))
	})

	log.Info("starting server", slog.String("address", cfg.Address))

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server", err)
	}

	log.Error("server stopped")
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
