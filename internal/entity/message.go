package entity

import "time"

type Message struct {
	Message  string
	Username string
	Times    time.Time
	UserId   string
	Pages    string
	Receive  string
}
