package entity

import "time"

var KeySpace = "chat-mqtt"

type Message struct {
	Topic     string
	Message   string
	Time      time.Time
	MessageId uint16
}
