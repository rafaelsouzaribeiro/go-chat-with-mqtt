package main

import (
	"log"
	"strconv"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/configs"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/infra/database/factory"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/infra/di"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/infra/web/mqtt/server"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/infra/web/webserver"
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

	portSocket, err := strconv.Atoi(Conf.PortWebsocketMqtt)
	if err != nil {
		log.Fatalf("Invalid port: %v", err)
	}

	webServer := webserver.NewWebServer(Conf.WebPort)
	webServer.AddHandlerChat(di)
	webServer.Start()

	svc := server.NewBroker(&server.Broker{
		Host:       Conf.HostMqtt,
		Port:       port,
		Username:   Conf.UserNameMqtt,
		Password:   Conf.PasswordMqtt,
		SocketHost: Conf.HostWebsocketMqtt,
		SocketPort: portSocket,
		Usecase:    di,
	})
	svc.StartServer()
}
