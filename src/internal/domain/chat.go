package domain

import "time"

type Message struct {
	Login string
	Time  time.Time
	Text  string
}
