package domain

import "time"

type Message struct {
	Login   string    `json:"login"`
	Time    time.Time `json:"time"`
	Message string    `json:"message"`
}
