package domain

type Lobby struct {
	ID       int64     `json:"id,omitempty"`
	LobbyURL string    `json:"lobby_url,omitempty"`
	VideoURL string    `json:"video_url,omitempty"`
	Timing   string    `json:"timing,omitempty"`
	Pause    bool      `json:"pause,omitempty"`
	Users    []*User   `json:"users,omitempty"`
	Chat     []Message `json:"chat,omitempty"`
}
