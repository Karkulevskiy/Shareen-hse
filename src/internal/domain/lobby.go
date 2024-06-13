package domain

// Lobby describes lobby.

type Lobby struct {
	ID       int64     `json:"id,omitempty"`
	LobbyURL string    `json:"lobby_url"`
	VideoURL string    `json:"video_url"`
	Timing   string    `json:"timing"`
	Pause    bool      `json:"pause"`
	Users    []*User   `json:"users"`
	Chat     []Message `json:"chat"`
}
