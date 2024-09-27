package main

import (
	"log"
	"strconv"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/configs"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/infra/database/factory"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/infra/di"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/infra/web/mqtt/server"
)

func main() {

	Conf, err := configs.LoadConfig("./")

	if err != nil {
		panic(err)
	}

	port, err := strconv.Atoi(Conf.PortMqtt)
	if err != nil {
		log.Fatalf("Invalid port: %v", err)
	}

	db, err := factory.NewFactory(&factory.Factory{
		Factory: factory.Cassandra,
	})

	if err != nil {
		panic(err)
	}

	di := di.NewUseCase(db)

	svc := server.NewBroker(Conf.HostMqtt, Conf.UserNameMqtt, Conf.PasswordMqtt, port, di)
	svc.StartServer()
}
