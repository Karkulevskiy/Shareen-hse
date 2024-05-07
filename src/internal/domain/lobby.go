package domain

type Lobby struct {
	ID       int64
	LobbyURL string
	VideoURL string
	Timing   string
	Pause    bool
	Users    []*User
	Chat     []Message
}
