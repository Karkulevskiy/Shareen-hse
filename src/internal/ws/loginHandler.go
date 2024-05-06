package ws

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/karkulevskiy/shareen/src/internal/lib"
	"github.com/karkulevskiy/shareen/src/internal/storage"
	"golang.org/x/crypto/bcrypt"
)

type Request struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Response struct {
	OTP     string `json:"otp,omitempty"`
	Message string `json:"message,omitempty"`
	Status  int    `json:"status,omitempty"`
}

func (m *Manager) LoginHandler(w http.ResponseWriter, r *http.Request) {
	const op = "ws.manager.loginHandler"

	log := m.log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var request Request

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Error("failed to decode login request", err)

		http.Error(w, "invalid login request", http.StatusBadRequest)
		w.Write(lib.Err("internal error", http.StatusInternalServerError))

		return
	}

	user, err := m.storage.User(request.Login)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			log.Info("user not found")

			http.Error(w, "user not found", http.StatusNotFound)
			w.Write(lib.Err("user not found", http.StatusNotFound))

			return
		}

		log.Error("failed to get user", err)

		w.Write(lib.Err("internal error", http.StatusInternalServerError))
		http.Error(w, "internal server error", http.StatusInternalServerError)

		return
	}

	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(request.Password)); err != nil {
		log.Info("invalid credentials", err)

		w.Write(lib.Err("invalid credentials", http.StatusUnauthorized))
		http.Error(w, "invalid credentials", http.StatusUnauthorized)

		return
	}

	response := Response{
		OTP: m.otps.NewOTP().Key,
	}

	data, err := json.Marshal(response)
	if err != nil {
		log.Error("failed to marshal response", err)

		w.Write(lib.Err("internal error", http.StatusInternalServerError))
		http.Error(w, "internal server error", http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (m *Manager) RegisterUser(w http.ResponseWriter, r *http.Request) {
	const op = "ws.manager.registerUser"

	log := m.log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var request Request

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Error("failed to decode register request", err)

		w.Write(lib.Err("invalid register request", http.StatusInternalServerError))

		return
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Info("failed to generate password hash", err)

		w.Write(lib.Err("internal error", http.StatusInternalServerError))

		return
	}

	err = m.storage.SaveUser(request.Login, passHash)
	if err != nil {
		if errors.Is(err, storage.ErrUserAlreadyExists) {
			log.Info("user already exists")

			w.Write(lib.Err("user already exists", http.StatusBadRequest))

			return
		}

		log.Error("failed to save user", err)

		w.Write(lib.Err("internal error", http.StatusInternalServerError))

		return
	}

	log.Info("user registered")

	resp := Response{
		Status:  http.StatusOK,
		Message: "user registered",
	}

	data, err := json.Marshal(&resp)
	if err != nil {
		log.Error("failed to marshal response", err)

		w.Write(lib.Err("internal error", http.StatusInternalServerError))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data))
}
