package domain

import "time"

// Message describes message.

type Message struct {
	Login   string    `json:"login"`
	Time    time.Time `json:"time"`
	Message string    `json:"message"`
}
