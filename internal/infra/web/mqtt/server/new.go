package server

import "github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase"

type Broker struct {
	Host       string
	Port       int
	Username   string
	Password   string
	SocketHost string
	SocketPort int
	Usecase    *usecase.UseCaseMessageUser
}

func NewBroker(b *Broker) *Broker {
	return b
}
