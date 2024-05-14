package domain

import "time"

type Message struct {
	Login string    `json:"login"`
	Time  time.Time `json:"time"`
	Text  string    `json:"text"`
}
