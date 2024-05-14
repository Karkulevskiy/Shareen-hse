package ws

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/karkulevskiy/shareen/src/internal/domain"
	"github.com/karkulevskiy/shareen/src/internal/lib"
	"github.com/karkulevskiy/shareen/src/internal/storage"
	"github.com/karkulevskiy/shareen/src/internal/ws/events"
)

// CreateLobbyHandler creates new lobby with unique URL
func CreateLobbyHandler(event Event, c *Client) {
	const op = "ws.CreateLobbyHandler"

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
	c.lobbyURL = lobbyURL

	payload, err := json.Marshal(events.CreateLobbyEventResponse{LobbyURL: lobbyURL})
	if err != nil {
		log.Error("failed to marshal lobby URL", err)
		SendResponseError(event.Type, http.StatusInternalServerError, c)
		return
	}

	response := CreateEvent(http.StatusOK, EventCreateLobby, payload)

	c.egress <- response
}

func JoinLobbyHandler(event Event, c *Client) {
	const op = "ws.JoinLobbyHandler"

	log := c.m.log.With(
		slog.String("op", op),
	)

	var request events.JoinLobbyEvent

	// Получаем json от клиента
	if err := json.Unmarshal(event.Payload, &request); err != nil {
		log.Error("failed to unmarshal join lobby request", err)
		SendResponseError(event.Type, http.StatusInternalServerError, c)
		return
	}

	// Ищем лобби по уникальному URL
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

	// Если в лобби, кто то есть еще, то обратимся к нему и спросим про актуальный тайминг в видео, и стоит ли пауза
	if len(c.m.lobbies[request.LobbyURL]) >= 1 {
		videoTimingCh := make(chan Event)

		c.m.videoTimingMap[request.Login] = videoTimingCh

		AskForVideoTiming(request.Login, c.m.lobbies[request.LobbyURL][0])

		responseTiming := <-videoTimingCh

		delete(c.m.videoTimingMap, request.Login)

		var respTimingData events.TimingEventResponse

		err := json.Unmarshal(responseTiming.Payload, &respTimingData)
		if err != nil {
			log.Error("failed to unmarshal response timing event", err)
			SendResponseError(event.Type, http.StatusInternalServerError, c)
			return
		}

		lobby.Pause = respTimingData.Pause
		lobby.Timing = respTimingData.Timing
	}

	// Добавим пользователя в кучу
	// Обновим для него логин, лобби
	c.login = request.Login
	c.m.lobbies[request.LobbyURL] = append(c.m.lobbies[request.LobbyURL], c)
	c.lobbyURL = request.LobbyURL

	for _, client := range c.m.lobbies[request.LobbyURL] {
		lobby.Users = append(lobby.Users, &domain.User{
			Login: client.login,
		})
	}

	// Создаем ответ
	payload, err := json.Marshal(&lobby)
	if err != nil {
		log.Error("failed to marshal lobby data", err)
		SendResponseError(event.Type, http.StatusInternalServerError, c)
		return
	}

	response := CreateEvent(http.StatusOK, EventJoinLobby, payload)

	// Отправляем ответ
	c.egress <- response

	log.Info("user joined lobby")

	// Уведомляем остальных, что подключился новый пользователь
	notifyUserJoinLobby(event, c, request)
}

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

	//send in chat that user was connected
	for _, client := range c.m.lobbies[request.LobbyURL] {
		if client != c {
			client.egress <- response
		}
	}
}

func notifyUserDisconnect(c *Client) {
	const op = "ws.notifyUserDisconnect"
	//TODO а как узнать лобби, если человек выходит из лобби
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

func InsertVideoHandler(event Event, c *Client) {
	const op = "ws.InsertVideoHandler"

	log := c.m.log.With(
		slog.String("op", op),
	)

	var insertReq events.InsertVideoEvent

	err := json.Unmarshal(event.Payload, &insertReq)
	if err != nil {
		log.Error("failed to unmarshal insert video request", err)
		SendResponseError(event.Type, http.StatusInternalServerError, c)
		return
	}

	err = c.m.storage.InsertVideo(insertReq.LobbyURL, insertReq.VideoURL)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, storage.ErrLobbyNotFound) {
			SendResponseError(event.Type, http.StatusBadRequest, c)
			return
		}
		log.Error("failed to insert video url", err)
		SendResponseError(event.Type, http.StatusInternalServerError, c)
		return
	}

	payload, _ := json.Marshal(events.InsertVideoResponse{URL: insertReq.VideoURL})

	response := CreateEvent(http.StatusOK, EventInsertVideoURL, payload)

	// Messages for lobby users to update video URL
	for _, client := range c.m.lobbies[insertReq.LobbyURL] {
		client.egress <- response
	}
}

func PauseVideoHandler(event Event, c *Client) {
	const op = "ws.PauseVideoHandler"

	log := c.m.log.With(
		slog.String("op", op),
	)

	var pauseReq events.PauseVideoEvent

	err := json.Unmarshal(event.Payload, &pauseReq)
	if err != nil {
		log.Error("failed to unmarshal pause video request", err)
		SendResponseError(event.Type, http.StatusInternalServerError, c)
		return
	}

	payload, _ := json.Marshal(events.PauseVideoResponse{Pause: pauseReq.Pause})

	response := CreateEvent(http.StatusOK, EventPauseVideo, payload)

	// Messages for lobby users to START / PAUSE video
	for _, client := range c.m.lobbies[pauseReq.LobbyURL] {
		client.egress <- response
	}
}

func AskForVideoTiming(login string, c *Client) {
	const op = "ws.AskForVideoTimingHandler"

	log := c.m.log.With(
		slog.String("op", op),
	)

	log.Info("ask for video timing")

	payload, _ := json.Marshal(events.AskVideoTimingEvent{Login: login})

	response := CreateEvent(http.StatusOK, EventGetVideoTiming, payload)

	c.egress <- response
}

// Сюда отправляем ответ от Клиента с информацией о видео
func GetVideoTimingHandler(event Event, c *Client) {
	const op = "ws.GetVideoTiming"

	log := c.m.log.With(
		slog.String("op", op),
	)

	var videoTimingRequest events.GetVideoTimingEvent

	err := json.Unmarshal(event.Payload, &videoTimingRequest)
	if err != nil {
		c.m.videoTimingMap[videoTimingRequest.Login] <- Event{Status: http.StatusInternalServerError}

		log.Error("failed to unmarshal video timing request", err)
		SendResponseError(event.Type, http.StatusInternalServerError, c)
		return
	}

	log.Info("get video timing", slog.String("user_login", videoTimingRequest.Login))

	payload, _ := json.Marshal(events.GetVideoTimingResponse{
		Timing: videoTimingRequest.Timing,
		Pause:  videoTimingRequest.Pause,
	})

	response := CreateEvent(http.StatusOK, EventGetVideoTiming, payload)

	c.m.videoTimingMap[videoTimingRequest.Login] <- response
}

// RewindVideoHandler rewinds video in lobby
func RewindVideoHandler(event Event, c *Client) {
	const op = "ws.RewindVideoHandler"

	log := c.m.log.With(
		slog.String("op", op),
	)

	var rewindReq events.RewindVideoEvent

	err := json.Unmarshal(event.Payload, &rewindReq)
	if err != nil {
		log.Error("failed to unmarshal rewind video request", err)
		SendResponseError(event.Type, http.StatusInternalServerError, c)
		return
	}

	for _, client := range c.m.lobbies[rewindReq.LobbyURL] {
		client.egress <- event
	}
}
