package domain

import "time"

type Lobby struct {
	ID       int64
	LobbyURL string
	VideoURL string
	Timing   time.Time
	Pause    bool
	Users    []*User
	Chat     []Message
}
