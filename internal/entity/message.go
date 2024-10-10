package entity

import "time"

type Message struct {
	Id        string
	Message   string
	Username  string
	Times     time.Time
	UserId    string
	Pages     string
	Receive   string
	Types     string
	PageTotal int64
}
