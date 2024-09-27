package server

import "github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase"

type Broker struct {
	host     string
	port     int
	username string
	password string
	Usecase  *usecase.UseCaseMessageUser
}

func NewBroker(host, username, password string, port int, usecase *usecase.UseCaseMessageUser) *Broker {
	return &Broker{
		host:     host,
		port:     port,
		username: username,
		password: password,
		Usecase:  usecase,
	}
}
