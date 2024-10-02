package dto

type PayloadMesage struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	UserId   string `json:"userId"`
}

type PayloadUser struct {
	Username string `json:"username"`
	Id       string `json:"id"`
	Photo    string `json:"photo"`
}
