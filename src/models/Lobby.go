package models

import "github.com/google/uuid"

type Lobby struct {
	ID       uuid.UUID
	LobbyURL string
	VideoURL string
}
