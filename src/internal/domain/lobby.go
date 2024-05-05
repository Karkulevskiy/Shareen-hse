package domain

type Lobby struct {
	ID       int64
	LobbyURL string
	VideoURL string
	Users    []*User
}
