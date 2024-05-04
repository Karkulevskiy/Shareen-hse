package storage

import "errors"

var (
	ErrURLExists     = errors.New("url already exists")
	ErrLobbyNotFound = errors.New("lobby not found")
	ErrUserNotFound  = errors.New("user not found")
)
