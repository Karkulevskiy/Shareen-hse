package ws

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/karkulevskiy/shareen/src/internal/lib"
	"github.com/karkulevskiy/shareen/src/internal/storage"
)

type JoinLobbyEvent struct {
	Login    string `json:"login"`
	LobbyURL string `json:"lobby_url"`
}

type LobbyDataEvent struct {
	Users    []string `json:"users"`
	Chat     []string `json:"chat"`
	VideoURL string   `json:"video_url"`
	Timings  string   `json:"timings"`
}

// CreateLobbyHandler creates new lobby with unique URL
func CreateLobbyHandler(event Event, c *Client) {
	const op = "ws.CreateLobbyHandler"

	type CreateResponse struct {
		LobbyURL string `json:"lobby_url"`
	}

	log := c.m.log.With(
		slog.String("op", op),
	)

	// Generate unique lobby URL
	lobbyURL := lib.GenerateURL()

	// Save lobby URL in DB
	lobbyURL, err := c.m.storage.CreateLobby(lobbyURL)
	if err != nil {
		log.Error("failed to create lobby", err)

		SendResponseError(event.Type, http.StatusInternalServerError, c)

		return
	}

	// Add client to lobby in RAM
	c.m.lobbies[lobbyURL] = append(c.m.lobbies[lobbyURL], c)

	data, err := json.Marshal(CreateResponse{LobbyURL: lobbyURL})
	if err != nil {
		log.Error("failed to marshal lobby URL", err)

		SendResponseError(event.Type, http.StatusInternalServerError, c)

		return
	}

	response := Event{
		Type:    EventCreateLobby,
		Status:  http.StatusOK,
		Payload: data,
	}

	c.egress <- response
}

func JoinLobbyHandler(event Event, c *Client) {
	const op = "ws.JoinLobbyHandler"

	log := c.m.log.With(
		slog.String("op", op),
	)

	var request JoinLobbyEvent

	if err := json.Unmarshal(event.Payload, &request); err != nil {
		log.Error("failed to unmarshal join lobby request", err)

		SendResponseError(event.Type, http.StatusInternalServerError, c)

		return
	}

	lobby, err := c.m.storage.Lobby(request.LobbyURL)
	if err != nil {
		fmt.Println("aa")
		if errors.Is(err, storage.ErrLobbyNotFound) {
			log.Info("lobby not found", err)

			SendResponseError(event.Type, http.StatusBadRequest, c)

			return

		}

		SendResponseError(event.Type, http.StatusInternalServerError, c)

		return
	}

	if _, ok := c.m.lobbies[request.LobbyURL]; !ok {
		log.Info("no lobby in RAM")
	}

	// Add user to lobby in RAM
	if _, ok := c.m.lobbies[request.LobbyURL]; ok {
		flag := false
		for _, client := range c.m.lobbies[request.LobbyURL] {
			if client == c {
				flag = true
				break
			}
		}
		if !flag {
			c.m.lobbies[request.LobbyURL] = append(c.m.lobbies[request.LobbyURL], c)
		}
	} else {
		c.m.lobbies[request.LobbyURL] = append(c.m.lobbies[request.LobbyURL], c)
	}

	//TODO: send notify in chat that user was disconnected
	response := JoinLobbyEvent{
		Login:    request.Login,
		LobbyURL: request.LobbyURL,
	}

	data, err := json.Marshal(&response)
	if err != nil {
		log.Error("failed to marshal join lobby response", err)

		SendResponseError(event.Type, http.StatusInternalServerError, c)

		return
	}

	userConnectedEvent := Event{
		Status:  http.StatusOK,
		Type:    EventUserJoinLobby,
		Payload: data,
	}

	//send in chat that user was connected
	for _, client := range c.m.lobbies[request.LobbyURL] {
		if client != c {
			client.egress <- userConnectedEvent
		}
	}

	lobbyData, err := json.Marshal(&lobby)
	if err != nil {
		log.Error("failed to marshal lobby data", err)

		SendResponseError(event.Type, http.StatusInternalServerError, c)

		return
	}

	lobbyResp := Event{
		Status:  http.StatusOK,
		Type:    EventJoinLobby,
		Payload: lobbyData,
	}

	c.egress <- lobbyResp

	log.Info("user joined lobby")
}

func InsertVideoHandler(event Event, c *Client) error {
	return nil
}

func PauseVideoHandler(event Event, c *Client) error {
	return nil
}

func StartVideoHandler(event Event, c *Client) error {
	return nil
}
