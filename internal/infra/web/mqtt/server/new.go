package server

import (
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase"
)

type Broker struct {
	Host     string
	Port     int
	Username string
	Password string
	Server   *mqtt.Server
	Topic    string
	Usecase  *usecase.UseCaseMessageUser
}

func NewBroker(b *Broker) *Broker {
	return b
}
