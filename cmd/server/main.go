package main

import (
	"log"
	"strconv"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/configs"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/infra/web/mqtt/server"
)

func main() {

	Conf, err := configs.LoadConfig("../")

	if err != nil {
		panic(err)
	}

	port, err := strconv.Atoi(Conf.PortMqtt)
	if err != nil {
		log.Fatalf("Invalid port: %v", err)
	}

	svc := server.NewBroker(Conf.HostMqtt, Conf.UserNameMqtt, Conf.PasswordMqtt, port)
	svc.StartServer()
}
