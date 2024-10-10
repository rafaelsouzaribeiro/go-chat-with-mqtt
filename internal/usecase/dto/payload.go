package dto

import "time"

type PayloadMesage struct {
	Username string    `json:"username"`
	Message  string    `json:"message"`
	UserId   string    `json:"userId"`
	Times    time.Time `json:"times"`
	Pages    string    `json:"pages"`
	Receive  string    `json:"receive"`
	Types    string    `json:"types"`
}

type PayloadUser struct {
	Username  string    `json:"username"`
	Id        string    `json:"id"`
	Photo     string    `json:"photo"`
	Times     time.Time `json:"times"`
	Password  string    `json:"password"`
	PageTotal int64     `json:"page_total"`
}
