package entity

import "time"

type User struct {
	Username  string    `json:"username"`
	Id        string    `json:"id"`
	Photo     string    `json:"photo"`
	Times     time.Time `json:"times"`
	Pages     string    `json:"page"`
	Password  string    `json:"password"`
	PageTotal int64     `json:"page_total"`
	Status    string    `json:"status"`
}
