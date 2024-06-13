package ws

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/karkulevskiy/shareen/src/internal/storage"
	"github.com/karkulevskiy/shareen/src/internal/ws/events"
)

// InsertVideoHandler inserts video in lobby.
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

// PauseVideoHandler pauses video in lobby.
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

// AskForVideoTimingHandler asks for video timing.
func AskForVideoTiming(login string, c *Client) {
	const op = "ws.AskForVideoTimingHandler"

	log := c.m.log.With(
		slog.String("op", op),
	)

	log.Info("ask for video timing")

	payload, err := json.Marshal(events.AskVideoTimingEvent{Login: login})
	if err != nil {
		log.Error("failed to unmarshal ask video timing request", err)
		SendResponseError(EventGetVideoTiming, http.StatusInternalServerError, c)
		return
	}

	response := CreateEvent(http.StatusOK, EventGetVideoTiming, payload)

	c.egress <- response
}

// GetVideoTimingHandler gets video timing.
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

	if _, ok := c.m.videoTimingMap[videoTimingRequest.Login]; ok {
		c.m.videoTimingMap[videoTimingRequest.Login] <- response
	}
}

// RewindVideoHandler rewinds video in lobby.
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

	payload, err := json.Marshal(&events.RewindVideoResponse{
		Timing:   rewindReq.Timing,
		LobbyURL: rewindReq.LobbyURL,
	})

	if err != nil {
		log.Error("failed to marshal rewind video response", err)
		SendResponseError(event.Type, http.StatusInternalServerError, c)
		return
	}

	response := CreateEvent(http.StatusOK, EventRewindVideo, payload)

	for _, client := range c.m.lobbies[rewindReq.LobbyURL] {
		client.egress <- response
	}
}
