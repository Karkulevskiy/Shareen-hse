package events

import "time"

type JoinLobbyEvent struct {
	Login    string `json:"login"`
	LobbyURL string `json:"lobby_url"`
}

type CreateLobbyEvent struct {
	Login string `json:"login"`
}

type CreateLobbyEventResponse struct {
	LobbyURL string `json:"lobby_url"`
}

type LobbyDataEvent struct {
	Users    []string `json:"users"`
	Chat     []string `json:"chat"`
	VideoURL string   `json:"video_url"`
	Timings  string   `json:"timings"`
}

type TimingEventResponse struct {
	Login  string `json:"login"`
	Timing string `json:"timing"`
	Pause  bool   `json:"pause"`
}

type InsertVideoEvent struct {
	VideoURL string `json:"video_url"`
	LobbyURL string `json:"lobby_url"`
}

type InsertVideoResponse struct {
	URL string `json:"url"`
}

type PauseVideoEvent struct {
	LobbyURL string `json:"lobby_url"`
	Pause    bool   `json:"pause"`
}

type PauseVideoResponse struct {
	Pause bool `json:"pause"`
}

type AskVideoTimingEvent struct {
	Login string `json:"login"`
}

type GetVideoTimingEvent struct {
	Login  string `json:"login"`
	Timing string `json:"timing"`
	Pause  bool   `json:"pause"`
}

type GetVideoTimingResponse struct {
	Timing string `json:"timing"`
	Pause  bool   `json:"pause"`
}

type RewindVideoEvent struct {
	LobbyURL string `json:"lobby_url"`
	Timing   string `json:"timing"`
}

type RewindVideoResponse struct {
	LobbyURL string `json:"lobby_url"`
	Timing   string `json:"timing"`
}

type UserDisconnectedEvent struct {
	Login string `json:"login"`
}

type SendMessageRequest struct {
	LobbyURL string `json:"lobby_url"`
	Login    string `json:"login"`
	Message  string `json:"message"`
}

type SendMessageResponse struct {
	Login   string    `json:"login"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}
