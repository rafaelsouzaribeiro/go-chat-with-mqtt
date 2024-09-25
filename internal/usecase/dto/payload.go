package dto

type Payload struct {
	Username  string `json:"username"`
	Topic     string `json:"topic"`
	Message   string `json:"message"`
	MessageId uint16 `json:"message_id"`
}
