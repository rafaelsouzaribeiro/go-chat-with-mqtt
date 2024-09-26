package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/configs"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/infra/web/mqtt/client"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase/dto"
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

	cli := client.NewClient(&client.Broker{
		Broker:   Conf.HostMqtt,
		Port:     port,
		Topic:    "topic/test",
		Username: Conf.UserNameMqtt,
		Password: Conf.PasswordMqtt,
	})

	channel := make(chan dto.Payload)
	go cli.Connect(channel)

	for messages := range channel {
		fmt.Printf("Message: %s Topic: %s Message ID: %d Username: %s \n",
			messages.Message, messages.Topic, messages.MessageId, messages.Username)
	}

	select {}
}
