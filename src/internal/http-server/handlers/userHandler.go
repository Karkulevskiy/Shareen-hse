package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/shareen/src/internal/domain/models"
)

type User interface {
	CreateUser(name string) (int64, error)
	GetUser(userID int64) (*models.User, error)
	DeleteUser(userID int64) error
	JoinUserToLobby(lobbyID string, userID string) error
}

// @Summary      Create user
// @Description  Create user
// @Tags         User
// @Produce      json
// @Param        name    path    string  false  "Create user using his name"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /user/{name} [post]
func CreateUser(log *slog.Logger, u User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "http-server.handlers.CreateUser"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		name := chi.URLParam(r, "name")

		if name == "" {
			log.Info("Name is empty")

			render.JSON(w, r, models.Response{
				Message: "name can't be empty",
				Status:  http.StatusBadRequest,
			})

			return
		}

		userID, err := u.CreateUser(name)

		if err != nil {
			log.Error("failed to create user", err)

			render.JSON(w, r, models.Response{
				Message: "failed to create user",
				Status:  http.StatusInternalServerError,
			})

			return
		}

		render.JSON(w, r, userID)
	}
}

// @Summary      Get user
// @Description  Get user by user id
// @Tags         User
// @Produce      json
// @Param        id    path    string  false  "get user by id"
// @Success      200 {object} models.User
// @Failure      400
// @Failure      500
// @Router       /user/{id} [get]
func GetUser(log *slog.Logger, u User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "http-server.handlers.GetUser"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		userID := chi.URLParam(r, "id")

		fmt.Println(userID)
		if userID == "" {
			log.Info("User ID is empty")

			render.JSON(w, r, models.Response{
				Message: "user id is empty",
				Status:  http.StatusBadRequest,
			})

			return
		}

		id, err := strconv.ParseInt(userID, 10, 64)
		if err != nil {
			log.Error("failed to parse user id", err)

			render.JSON(w, r, models.Response{
				Message: "failed to parse user id",
				Status:  http.StatusInternalServerError,
			})

			return
		}
		user, err := u.GetUser(id)
		if err != nil {
			log.Error("failed to get user", err)

			render.JSON(w, r, models.Response{
				Message: "failed to get user",
				Status:  http.StatusInternalServerError,
			})

			return
		}

		render.JSON(w, r, user)
	}
}

// @Summary      Delete user
// @Description  Delete user by id
// @Tags         User
// @Produce      json
// @Param        id    path    integer  false  "Delete user by ID"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /user/{id} [delete]
func DeleteUser(log *slog.Logger, u User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "http-server.handlers.DeleteUser"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		userID := chi.URLParam(r, "id")

		if userID == "" {
			log.Info("User can't be empty")

			render.JSON(w, r, models.Response{
				Message: "user id can't be empty",
				Status:  http.StatusBadRequest,
			})

			return
		}

		fmt.Println(userID)
		id, err := strconv.ParseInt(userID, 10, 64)
		if err != nil {
			log.Error("failed to parse user id", err)

			render.JSON(w, r, models.Response{
				Message: "failed to parse user id",
				Status:  http.StatusInternalServerError,
			})

			return
		}

		err = u.DeleteUser(id)
		if err != nil {
			log.Error("failed to delete user", err)

			render.JSON(w, r, models.Response{
				Message: "failed to delete user",
				Status:  http.StatusInternalServerError,
			})

			return
		}

		render.JSON(w, r, http.StatusOK)
	}
}

// @Summary      Join user
// @Description  Join user in lobby by userID, lobbyID
// @Tags         User
// @Produce      json
// @Param        url    path    string  false  "Join user in lobby by unique url"
// @Param        id    path    string  false  "Join user in lobby by unique url"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /user/{url}-{id} [patch]
func JoinUserToLobby(log *slog.Logger, u User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "http-server.handlers.JoinUserToLobby"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		lobbyID, userID := chi.URLParam(r, "url"), chi.URLParam(r, "id")

		if lobbyID == "" {
			log.Info("Lobby ID is empty")

			render.JSON(w, r, models.Response{
				Message: "lobby id is empty",
				Status:  http.StatusBadRequest,
			})

			return
		}

		if userID == "" {
			log.Info("User ID is empty")

			render.JSON(w, r, models.Response{
				Message: "user id is empty",
				Status:  http.StatusBadRequest,
			})

			return
		}

		err := u.JoinUserToLobby(lobbyID, userID)
		if err != nil {
			log.Error("failed to join user to lobby", err)

			render.JSON(w, r, models.Response{
				Message: "failed to join user to lobby",
				Status:  http.StatusInternalServerError,
			})

			return
		}

		render.JSON(w, r, http.StatusOK)

	}
}
