package entity

import "time"

var KeySpace = "chatmqtt"

type Message struct {
	Topic     string
	Message   string
	Username  string
	Time      time.Time
	MessageId uint16
}
