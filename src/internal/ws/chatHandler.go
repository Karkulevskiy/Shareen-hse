package ws

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

func SendMessageHandler(event Event, c *Client) {
	const op = "ws.SendMessage"

	log := c.m.log.With(
		slog.String("op", op),
	)

	type SendMessageRequest struct {
		LobbyURL string `json:"lobby_url"`
		Login    string `json:"login"`
		Message  string `json:"message"`
	}

	var req SendMessageRequest
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

	type SendMessageResponse struct {
		Login   string    `json:"login"`
		Message string    `json:"message"`
		Time    time.Time `json:"time"`
	}

	sendData, _ := json.Marshal(SendMessageResponse{
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
