package server

import (
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase"
)

type Broker struct {
	host     string
	port     int
	username string
	password string
	Server   *mqtt.Server
	topic    string
	Usecase  *usecase.UseCaseMessageUser
}

func NewBroker(host, username, password, topic string, port int, Usecase *usecase.UseCaseMessageUser) *Broker {
	return &Broker{
		host:     host,
		port:     port,
		username: username,
		password: password,
		topic:    topic,
		Usecase:  Usecase,
	}
}
