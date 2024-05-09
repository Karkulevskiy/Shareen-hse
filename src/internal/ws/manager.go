package ws

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/websocket"
	"github.com/karkulevskiy/shareen/src/internal/storage/postgres"
)

var (
	websocketUpgrader = websocket.Upgrader{
		//TODO: CheckOrigin:     checkOrigin, НАТСРОИТЬ CORS
		ReadBufferSize:  1024, //TODO: посмотреть сколько нужно ставить буффер
		WriteBufferSize: 1024,
		CheckOrigin:     checkOrigin,
	}
)

type Manager struct {
	log *slog.Logger
	sync.RWMutex
	handlers       map[string]EventHandler
	storage        *postgres.Postgres
	otps           RetentionMap
	clients        map[*Client]bool
	lobbies        map[string][]*Client
	videoTimingMap map[string]chan Event
}

func NewManager(storage *postgres.Postgres, log *slog.Logger, ctx context.Context) *Manager {
	m := &Manager{
		handlers:       make(map[string]EventHandler),
		storage:        storage,
		log:            log,
		otps:           NewRetentionMap(ctx, 5*time.Minute), //TODO: потом выбрать время действия OTP
		clients:        make(map[*Client]bool),
		lobbies:        make(map[string][]*Client),
		videoTimingMap: make(map[string]chan Event),
	}
	m.setupEventHandlers()

	return m
}

func (m *Manager) setupEventHandlers() {
	m.handlers[EventCreateLobby] = CreateLobbyHandler
	m.handlers[EventJoinLobby] = JoinLobbyHandler
	m.handlers[EventInsertVideoURL] = InsertVideoHandler
	m.handlers[EventPauseVideo] = PauseVideoHandler
	m.handlers[EventSendMessage] = SendMessageHandler
	m.handlers[EventGetVideoTiming] = GetVideoTimingHandler
	m.handlers[EventRewindVideo] = RewindVideoHandler
	// m.handlers[]
}

func (m *Manager) routeEvent(event Event, c *Client) {
	if handler, ok := m.handlers[event.Type]; ok {
		handler(event, c)
	} else {
		m.log.Error("no such event type")
	}
}

func (m *Manager) ServeWS(w http.ResponseWriter, r *http.Request) {
	const op = "ws.manager.serveWS"

	type Response struct {
		Message string `json:"message"`
		Status  int    `json:"status"`
	}

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

	if !m.otps.VerifyOTP(otp) {
		log.Info("user is not authorized")

		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	log.Info("new connection")

	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error("failed to upgrade connection", err)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	client := NewClient(conn, m)
	m.addClient(client)

	go client.readMessages()
	go client.writeMessages()
}

func (m *Manager) addClient(c *Client) {
	m.Lock()
	defer m.Unlock()
	m.clients[c] = true
}

func (m *Manager) removeClient(c *Client) {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.clients[c]; ok {
		c.conn.Close()
		delete(m.clients, c)
	}
}

func (m *Manager) clientInLobby(lobbyURL string, c *Client) bool {
	m.Lock()
	defer m.Unlock()
	for _, client := range m.lobbies[lobbyURL] {
		if client == c {
			return true
		}
	}
	return false
}

func checkOrigin(r *http.Request) bool {

	// // Grab the request origin
	// origin := r.Header.Get("Origin")

	// switch origin {
	// // Update this to HTTPS
	// case "http://localhost:5001":
	// 	return true
	// default:
	// 	return false
	// }
	return true
}

func accessControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		next.ServeHTTP(w, r)
	})
}
