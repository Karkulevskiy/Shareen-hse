package models

type Lobby struct {
	ID        int64
	VideoURL  string
	LobbdyURL string
	Users     []User
}
