package ws

import (
	"encoding/json"
)

const (
	// Events for lobby logic
	EventCreateLobby = "create_lobby"
	EventDisconnect  = "disconnect"
	EventJoinLobby   = "join_lobby"

	// Events for video logic
	EventPauseVideo     = "pause_video"
	EventStartVideo     = "start_video"
	EventRewindVideo    = "rewind_video"
	EventInsertVideoURL = "insert_video_url"
	EventGetVideoTiming = "get_video_timing"

	// Events for login, register logic
	EventLogin    = "login"
	EventRegister = "register"

	// Events for chat logic
	EventUserJoinLobby  = "user_join_lobby"
	EventUserLeaveLobby = "user_leave_lobby"
	EventSendMessage    = "send_message"
	EventNewMessage     = "new_message"

	EventLobbyNotFound = "lobby_not_found"
)

type EventHandler func(event Event, c *Client)

// Event struct
type Event struct {
	Type    string          `json:"type,omitempty"`
	Status  int             `json:"status,omitempty"`
	Payload json.RawMessage `json:"payload,omitempty"`
}

// SendResponseError sends error response.
func SendResponseError(eventType string, status int, c *Client) {

	type ResponseErr struct {
		Status int `json:"status"`
	}

	data, _ := json.Marshal(ResponseErr{Status: status})

	respEvent := Event{
		Status:  status,
		Type:    eventType,
		Payload: data,
	}

	c.egress <- respEvent
}

// CreateEvent creates new event.
func CreateEvent(status int, eventType string, payload json.RawMessage) Event {
	return Event{
		Status:  status,
		Type:    eventType,
		Payload: payload,
	}
}
