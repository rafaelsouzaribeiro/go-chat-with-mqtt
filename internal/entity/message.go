package entity

import "time"

var KeySpace = "chat-mqtt"

type Message struct {
	Topic     string
	Message   string
	Username  string
	Time      time.Time
	MessageId uint16
}
