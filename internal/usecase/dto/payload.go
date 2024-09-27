package dto

type Payload struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	UserId   string `json:"userId"`
}
