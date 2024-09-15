package entity

import "time"

var KeySpace = "whatsapp"

type Message struct {
	Message string
	Type    string
	Time    time.Time
}
