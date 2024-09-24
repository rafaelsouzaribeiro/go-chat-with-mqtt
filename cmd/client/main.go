package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/configs"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/infra/di"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/infra/web/mqtt/client"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase/dto"
)

func main() {
	Conf, err := configs.LoadConfig("../")

	if err != nil {
		panic(err)
	}

	di := di.NewUseCase()

	port, err := strconv.Atoi(Conf.PortMqtt)
	if err != nil {
		log.Fatalf("Invalid port: %v", err)
	}

	cli := client.NewClient(di, &client.Broker{
		Broker:   Conf.HostMqtt,
		Port:     port,
		Topic:    "topic/test",
		Username: Conf.UserNameMqtt,
		Password: Conf.PasswordMqtt,
	})

	channel := make(chan dto.Payload)
	cli.Connect(channel)

	err = cli.PublishMessage(&dto.Payload{
		Username: "Rafael",
		Message:  "Testar",
		Topic:    "topic/test",
	})

	if err != nil {
		panic(err)
	}
	for messages := range channel {
		fmt.Printf("Message: %s Topic: %s Message ID: %d \n",
			messages.Message, messages.Topic, messages.MessageId)

	}

	select {}
}
