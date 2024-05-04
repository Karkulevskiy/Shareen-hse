package ws

import "encoding/json"

const (
	// Events for lobby logic
	EventCreateLobby = "create_lobby"
	EventJoinLobby   = "join_lobby"

	// Events for video logic
	EventPauseVideo     = "pause_video"
	EventStartVideo     = "start_video"
	EventRewindVideo    = "rewind_video"
	EventInsertVideoURL = "insert_video_url"

	// Events for chat logic
	EventSendMessage = "send_message"
	EventNewMessage  = "new_message"
)

type EventHandler func(event Event, c *Client) error

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}
