package ws

import (
	"encoding/json"
	"log/slog"
	"time"

	"github.com/gorilla/websocket"
)

var (
	pongWait   = time.Minute * 10
	pingPeriod = (pongWait * 9) / 10
)

type Client struct {
	conn   *websocket.Conn
	m      *Manager
	egress chan Event
}

func NewClient(conn *websocket.Conn, m *Manager, login string) *Client {
	return &Client{
		conn:   conn,
		m:      m,
		egress: make(chan Event),
	}
}

func (c *Client) readMessages() {
	defer func() {
		c.m.removeClient(c)
	}()

	const op = "ws.client.readMessages"

	log := c.m.log.With(
		slog.String("op", op),
	)

	if err := c.conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Info("failed to set read deadline", err)

		return
	}

	c.conn.SetReadLimit(512)
	c.conn.SetPongHandler(c.pongHandler)

	for {
		_, p, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Error("failed to read message", err)
			}

			break
		}

		var request Event

		if err := json.Unmarshal(p, &request); err != nil {
			log.Error("failed to unmarshal request", err)

			break
		}

		c.m.routeEvent(request, c)
	}
}

func (c *Client) writeMessages() {
	defer func() {
		c.m.removeClient(c)
	}()

	const op = "ws.client.writeMessages"

	log := c.m.log.With(
		slog.String("op", op),
	)

	ticker := time.NewTicker(pingPeriod)

	for {
		select {
		case msg, ok := <-c.egress:
			if !ok {
				if err := c.conn.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Error("failed to close connection", err)
				}

				return
			}

			data, err := json.Marshal(msg)
			if err != nil {
				log.Error("failed to marshal message", err)

				return
			}

			if err := c.conn.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Error("failed to write message", err)
			}

			log.Info("message sent", slog.String("message", string(data)))

		case <-ticker.C:
			log.Info("ping")

			if err := c.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				log.Error("failed to ping", err)

				return
			}
		}
	}
}

func (c *Client) pongHandler(pongMsg string) error {
	c.m.log.Info("pong")

	return c.conn.SetReadDeadline(time.Now().Add(pongWait))
}
