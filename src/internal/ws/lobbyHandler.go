package ws

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/karkulevskiy/shareen/src/internal/domain"
	"github.com/karkulevskiy/shareen/src/internal/lib"
	"github.com/karkulevskiy/shareen/src/internal/storage"
	"github.com/karkulevskiy/shareen/src/internal/ws/events"
)

// CreateLobbyHandler creates new lobby with unique URL.
func CreateLobbyHandler(event Event, c *Client) {
	const op = "ws.CreateLobbyHandler"

	log := c.m.log.With(
		slog.String("op", op),
	)

	var createLobbyEvent events.CreateLobbyEvent

	err := json.Unmarshal(event.Payload, &createLobbyEvent)
	if err != nil {
		log.Error("failed to unmarshal create lobby request", err)
		SendResponseError(event.Type, http.StatusInternalServerError, c)
		return
	}

	lobbyURL := lib.GenerateURL()

	_, err = c.m.storage.CreateLobby(lobbyURL)
	if err != nil {
		log.Error("failed to create lobby", err)
		SendResponseError(event.Type, http.StatusInternalServerError, c)
		return
	}

	c.m.lobbies[lobbyURL] = append(c.m.lobbies[lobbyURL], c)
	c.lobbyURL = lobbyURL
	c.login = createLobbyEvent.Login

	payload, err := json.Marshal(events.CreateLobbyEventResponse{LobbyURL: lobbyURL})
	if err != nil {
		log.Error("failed to marshal lobby URL", err)
		SendResponseError(event.Type, http.StatusInternalServerError, c)
		return
	}

	response := CreateEvent(http.StatusOK, EventCreateLobby, payload)

	c.egress <- response
}

// JoinLobbyHandler adds user to lobby.
func JoinLobbyHandler(event Event, c *Client) {
	const op = "ws.JoinLobbyHandler"

	log := c.m.log.With(
		slog.String("op", op),
	)

	var request events.JoinLobbyEvent

	if err := json.Unmarshal(event.Payload, &request); err != nil {
		log.Error("failed to unmarshal join lobby request", err)
		SendResponseError(event.Type, http.StatusInternalServerError, c)
		return
	}

	lobby, err := c.m.storage.Lobby(request.LobbyURL)
	if err != nil {
		if errors.Is(err, storage.ErrLobbyNotFound) {
			log.Info("lobby not found", err)
			SendResponseError(event.Type, http.StatusBadRequest, c)
			return
		}

		log.Error("failed to get lobby", err)
		SendResponseError(event.Type, http.StatusInternalServerError, c)
		return
	}

	if isUserInLobby(request.Login, lobby,
		request, c.m.lobbies[request.LobbyURL], c, event) {
		return
	}

	respTiming := tryGetTiming(request, c)

	if respTiming != nil {
		var respTimingData events.TimingEventResponse

		err := json.Unmarshal(respTiming.Payload, &respTimingData)
		if err != nil {
			log.Error("failed to unmarshal response timing event", err)
			SendResponseError(event.Type, http.StatusInternalServerError, c)
			return
		}

		lobby.Pause = respTimingData.Pause
		lobby.Timing = respTimingData.Timing

	}

	c.login = request.Login
	c.m.lobbies[request.LobbyURL] = append(c.m.lobbies[request.LobbyURL], c)
	c.lobbyURL = request.LobbyURL

	for _, client := range c.m.lobbies[request.LobbyURL] {
		lobby.Users = append(lobby.Users, &domain.User{
			Login: client.login,
		})
	}

	payload, err := json.Marshal(&lobby)
	if err != nil {
		log.Error("failed to marshal lobby data", err)
		SendResponseError(event.Type, http.StatusInternalServerError, c)
		return
	}

	response := CreateEvent(http.StatusOK, EventJoinLobby, payload)

	c.egress <- response

	log.Info("user joined lobby")

	notifyUserJoinLobby(event, c, request)
}

// tryGetTiming try's to get video timing from users.
func tryGetTiming(request events.JoinLobbyEvent, c *Client) *Event {
	videoTimingCh := make(chan Event)
	var respTiming Event
	for _, client := range c.m.lobbies[request.LobbyURL] {
		if client.login != request.Login {
			c.m.videoTimingMap[request.Login] = videoTimingCh

			AskForVideoTiming(request.Login, c.m.lobbies[request.LobbyURL][0])

			select {
			case respTiming = <-videoTimingCh:
				break
			case <-time.After(5 * time.Second):
				c.m.log.Info("no timing response")
				delete(c.m.videoTimingMap, request.Login)
			}
		}
	}

	return &respTiming
}

// notifyUserJoinLobby notifies users that user joined lobby.
func notifyUserJoinLobby(event Event, c *Client, request events.JoinLobbyEvent) {
	const op = "ws.userJoinLobby"

	log := c.m.log.With(
		slog.String("op", op),
	)

	payload, err := json.Marshal(events.JoinLobbyEvent{
		Login:    request.Login,
		LobbyURL: request.LobbyURL,
	})

	if err != nil {
		log.Error("failed to marshal join lobby response", err)
		SendResponseError(event.Type, http.StatusInternalServerError, c)
		return
	}

	response := CreateEvent(http.StatusOK, EventUserJoinLobby, payload)

	log.Info("notify user join lobby")

	for _, client := range c.m.lobbies[request.LobbyURL] {
		if client != c {
			client.egress <- response
		}
	}
}

// notifyUserDisconnect notifies users that user disconnected.
func notifyUserDisconnect(c *Client) {
	const op = "ws.notifyUserDisconnect"
	log := c.m.log.With(
		slog.String("op", op),
	)

	if _, ok := c.m.lobbies[c.lobbyURL]; !ok {
		log.Info("no lobby found")
		return
	}

	lobby := c.m.lobbies[c.lobbyURL]
	for i, client := range lobby {
		if client == c {
			lobby = append(lobby[:i], lobby[i+1:]...)
			break
		}
	}

	c.m.lobbies[c.lobbyURL] = lobby

	log.Info("client was removed")

	payload, err := json.Marshal(events.UserDisconnectedEvent{Login: c.login})

	if err != nil {
		log.Error("failed to marshal user disconnected event", err)
		SendResponseError(EventDisconnect, http.StatusInternalServerError, c)
		return
	}

	event := CreateEvent(http.StatusOK, EventDisconnect, payload)

	for _, clients := range c.m.lobbies[c.lobbyURL] {
		clients.egress <- event
	}
}


// isUserInLobby checks if user already in lobby.
func isUserInLobby(login string, lobby *domain.Lobby,
	request events.JoinLobbyEvent,
	clients []*Client,
	c *Client, event Event) bool {
	for _, client := range clients {
		if client.login == login {
			c.m.log.Info("user already in lobby")

			for _, client := range c.m.lobbies[request.LobbyURL] {
				lobby.Users = append(lobby.Users, &domain.User{
					Login: client.login,
				})
			}

			payload, err := json.Marshal(&lobby)
			if err != nil {
				c.m.log.Error("failed to marshal lobby data", err)
				SendResponseError(event.Type, http.StatusInternalServerError, c)
				return true
			}

			response := CreateEvent(http.StatusOK, EventJoinLobby, payload)

			// Отправляем ответ
			c.egress <- response
			return true
		}
	}
	return false
}
