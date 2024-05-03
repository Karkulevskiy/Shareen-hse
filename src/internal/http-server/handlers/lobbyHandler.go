package handlers

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/shareen/src/internal/domain/models"
	"github.com/shareen/src/internal/storage"
	"github.com/shareen/src/internal/utils"
)

type Lobbier interface {
	CreateLobby(lobbyURL string) (string, error)
	GetLobby(lobbyURL string) (*models.Lobby, error)
	DeleteLobby(lobbyURL string) error
	SetLobbyVideoURL(lobbyID string, videoURL string) error
}

// @Summary      Get lobby
// @Description  Get lobby by unique url
// @Tags         Lobby
// @Produce      json
// @Param        url    path    string  false  "lobby search by url"
// @Success      200  {object}   models.Lobby
// @Failure      400
// @Failure      500
// @Router       /lobby/{url} [get]
func GetLobby(log *slog.Logger, l Lobbier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "http-server.handlers.GetLobby"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		lobbyURL := chi.URLParam(r, "url")

		if lobbyURL == "" {
			log.Info("Lobby URL is empty")

			render.JSON(w, r, models.Response{
				Message: "lobby url is empty",
				Status:  http.StatusBadRequest,
			})

			return
		}

		lobby, err := l.GetLobby(lobbyURL)

		if err != nil {
			if errors.Is(err, storage.ErrLobbyNotFound) {
				log.Info("lobby not found", "lobby url", lobbyURL)

				render.JSON(w, r, models.Response{
					Message: "lobby not found",
					Status:  http.StatusBadRequest,
				})

				return
			}

			log.Error("failed to get lobby", err)

			render.JSON(w, r, models.Response{
				Message: "failed to get lobby",
				Status:  http.StatusInternalServerError,
			})

			return
		}

		render.JSON(w, r, lobby)
	}
}

// @Summary      Create lobby
// @Description  Create lobby | ПОТОМ НУЖНО ПОЛУЧАТЬ ID ПОЛЬЗОВАТЕЛЯ, КОТОРЫЙ СОЗДАЕТ ЛОББИ
// @Tags         Lobby
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /lobby/ [post]
func CreateLobby(log *slog.Logger, l Lobbier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "http-server.handlers.CreateLobby"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		lobbyURL := utils.CreateURL()

		lobbyID, err := l.CreateLobby(lobbyURL)
		if err != nil {
			log.Error("failed to create lobby", err)

			render.JSON(w, r, models.Response{
				Message: "failed to create lobby",
				Status:  http.StatusInternalServerError,
			})

			return
		}

		fmt.Println(lobbyID)

		render.JSON(w, r, lobbyID)
	}
}

// @Summary      Delete lobby
// @Description  Delete lobby by unique url
// @Tags         Lobby
// @Produce      json
// @Param        url    path    string  false  "delete lobby by unique url"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /lobby/{url} [delete]
func DeleteLobby(log *slog.Logger, l Lobbier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "http-server.handlers.DeleteLobby"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		lobbyURL := chi.URLParam(r, "url")

		if lobbyURL == "" {
			log.Info("Lobby URL is empty")

			render.JSON(w, r, models.Response{
				Message: "lobby url is empty",
				Status:  http.StatusBadRequest,
			})

			return
		}

		err := l.DeleteLobby(lobbyURL)
		if err != nil {
			log.Error("failed to delete lobby", err)

			render.JSON(w, r, models.Response{
				Message: "failed to delete lobby",
				Status:  http.StatusInternalServerError,
			})

			return
		}

		render.JSON(w, r, http.StatusOK)
	}
}

// @Summary      Update vieo URL in lobby
// @Description  Update vieo URL in lobby
// @Tags         Lobby
// @Produce      json
// @Param        url    path    string  false  "update video URL in lobby by unique url"
// @Param        video    path    string  false  "update video URL in lobby by unique url"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /lobby/{url}-{video} [patch]
func SetLobbyVideoURL(log *slog.Logger, l Lobbier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "http-server.handlers.UpdateLobbyVideoURL"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		lobbyURL, videoURL := chi.URLParam(r, "url"), chi.URLParam(r, "video")

		if lobbyURL == "" {
			log.Info("Lobby URL is empty")

			render.JSON(w, r, models.Response{
				Message: "lobby url is empty",
				Status:  http.StatusBadRequest,
			})

			return
		}

		if videoURL == "" {
			log.Info("Video URL is empty")

			render.JSON(w, r, models.Response{
				Message: "video url is empty",
				Status:  http.StatusBadRequest,
			})

			return
		}

		iframe, err := utils.GetIframe(videoURL)
		if err != nil {
			log.Info("failed to get iframe", err)

			render.JSON(w, r, models.Response{
				Message: "failed to get iframe",
				Status:  http.StatusBadRequest,
			})

			return
		}

		err = l.SetLobbyVideoURL(lobbyURL, iframe)
		if err != nil {
			log.Error("failed to update lobby video url", err)

			render.JSON(w, r, models.Response{
				Message: "failed to update lobby video url",
				Status:  http.StatusInternalServerError,
			})

			return
		}

		render.JSON(w, r, iframe)
	}
}
