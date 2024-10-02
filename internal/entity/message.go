package entity

import "time"

type Message struct {
	Message  string
	Username string
	Time     time.Time
	UserId   string
}
