package repositories

import "errors"

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrLobbyNotFound = errors.New("lobby not found")
)
