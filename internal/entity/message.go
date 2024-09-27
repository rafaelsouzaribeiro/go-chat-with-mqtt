package entity

import "time"

var KeySpace = "chatmqtt"

type Message struct {
	Message  string
	Username string
	Time     time.Time
	UserId   string
}
