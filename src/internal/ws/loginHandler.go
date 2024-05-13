package ws

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/karkulevskiy/shareen/src/internal/lib"
	"github.com/karkulevskiy/shareen/src/internal/storage"
	"golang.org/x/crypto/bcrypt"
)

func (m *Manager) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	const op = "ws.manager.loginHandler"

	log := m.log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	type Request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	type ResponseOTP struct {
		OTP string `json:"otp"`
	}

	var request Request

	fmt.Println(r.Body)

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Error("failed to decode login request", err)

		lib.HTPPErr(w, EventLogin, http.StatusInternalServerError)

		return
	}

	user, err := m.storage.User(request.Login)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			log.Info("user not found")

			lib.HTPPErr(w, EventLogin, http.StatusBadRequest)

			return
		}

		log.Error("failed to get user", err)

		lib.HTPPErr(w, EventLogin, http.StatusInternalServerError)

		return
	}

	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(request.Password)); err != nil {
		log.Info("invalid credentials", err)

		lib.HTPPErr(w, EventLogin, http.StatusUnauthorized)

		return
	}

	otp := ResponseOTP{
		OTP: m.otps.NewOTP().Key,
	}

	otpJSON, _ := json.Marshal(otp)

	response := Event{
		Type:    EventLogin,
		Status:  http.StatusOK,
		Payload: otpJSON,
	}

	data, err := json.Marshal(response)
	if err != nil {
		log.Error("failed to marshal response", err)

		lib.HTPPErr(w, EventLogin, http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (m *Manager) RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	const op = "ws.manager.registerUser"

	log := m.log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	type Request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	var request Request

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Error("failed to decode register request", err)
		lib.HTPPErr(w, EventLogin, http.StatusInternalServerError)
		return
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Info("failed to generate password hash", err)
		lib.HTPPErr(w, EventLogin, http.StatusInternalServerError)
		return
	}

	err = m.storage.SaveUser(request.Login, passHash)
	if err != nil {
		if errors.Is(err, storage.ErrUserAlreadyExists) {
			log.Info("user already exists")
			lib.HTPPErr(w, EventLogin, http.StatusBadRequest)
			return
		}

		log.Error("failed to save user", err)
		lib.HTPPErr(w, EventLogin, http.StatusInternalServerError)
		return
	}

	log.Info("user registered")

	resp := Event{
		Type:    EventRegister,
		Status:  http.StatusOK,
		Payload: nil,
	}

	data, err := json.Marshal(&resp)
	if err != nil {
		log.Error("failed to marshal response", err)
		lib.HTPPErr(w, EventLogin, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data))
}
