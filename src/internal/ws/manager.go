package ws

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"sync"

	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/websocket"
	"github.com/karkulevskiy/shareen/src/internal/storage/postgres"
)

var (
	websocketUpgrader = websocket.Upgrader{
		//TODO: CheckOrigin:     checkOrigin, НАТСРОИТЬ CORS
		ReadBufferSize:  1024, //TODO: посмотреть сколько нужно ставить буффер
		WriteBufferSize: 1024,
	}
)

type Manager struct {
	log *slog.Logger
	sync.RWMutex
	handlers map[string]EventHandler
	storage  *postgres.Storage
}

func NewManager(storage *postgres.Storage, log *slog.Logger) *Manager {
	m := &Manager{
		handlers: make(map[string]EventHandler),
		storage:  storage,
		log:      log,
	}

	m.setupEventHandlers()

	return m
}

func (m *Manager) setupEventHandlers() {
	m.handlers[EventCreateLobby] = handlers.CreateLobby
	//TODO:
	// m.handlers[]
}

func (m *Manager) routeEvent(event Event, c *Client) error {
	if handler, ok := m.handlers[event.Type]; ok {
		if err := handler(event, c); err != nil {
			m.log.Error("failed to handle event", err)
			return err
		}
		return nil
	} else {
		return errors.New("no such event type")
	}
}

func (m *Manager) serveWS(w http.ResponseWriter, r *http.Request) {
	const op = "ws.manager.serveWS"

	log := m.log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	otp := r.URL.Query().Get("otp")
	if otp == "" {
		log.Info("otp is empty")

		resp, err := json.Marshal(Response{
			Message: "otp is empty",
			Status:  400,
		})

		if err != nil {
			log.Error("failed to marshal response", err)

			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write(resp)
	}

	log.Info("new connection")

}
