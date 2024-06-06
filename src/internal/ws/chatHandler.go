package ws

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/karkulevskiy/shareen/src/internal/ws/events"
)

// SendMessageHandler sends message in certail lobby.
func SendMessageHandler(event Event, c *Client) {
	const op = "ws.SendMessage"

	log := c.m.log.With(
		slog.String("op", op),
	)

	var req events.SendMessageRequest

	if err := json.Unmarshal(event.Payload, &req); err != nil {
		log.Error("failed to unmarshal send message request", err)
		SendResponseError(event.Type, http.StatusInternalServerError, c)
		return
	}

	err := c.m.storage.SaveMessage(req.LobbyURL, req.Login, req.Message)
	if err != nil {
		log.Error("failed to save message", err)
		SendResponseError(event.Type, http.StatusInternalServerError, c)
		return
	}

	sendData, _ := json.Marshal(events.SendMessageResponse{
		Login:   req.Login,
		Message: req.Message,
		Time:    time.Now(),
	})

	for _, client := range c.m.lobbies[req.LobbyURL] {
		client.egress <- Event{
			Status:  http.StatusOK,
			Type:    EventSendMessage,
			Payload: sendData,
		}
	}
}
