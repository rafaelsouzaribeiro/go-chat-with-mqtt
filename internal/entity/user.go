package entity

import "time"

type User struct {
	Username string
	Id       string
	Photo    string
	Times    time.Time
	Pages    string
}
