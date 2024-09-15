package dto

import "time"

type MessageDto struct {
	Message string
	Type    string
	Time    time.Time
	Id      string
}
