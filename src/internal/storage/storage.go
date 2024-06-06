package storage

import "errors"

var (
	ErrURLExists     = errors.New("url already exists")
	ErrLobbyNotFound = errors.New("lobby not found")
	ErrUserNotFound  = errors.New("user not found")

	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserAlreadyExists  = errors.New("user already exists")
)
