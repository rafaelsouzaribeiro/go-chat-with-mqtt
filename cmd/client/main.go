package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/configs"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/infra/database/cassandra/connection"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/infra/di"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/infra/web/mqtt/client"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase/dto"
)

func main() {
	Conf, err := configs.LoadConfig("../")

	if err != nil {
		panic(err)
	}

	hosts := strings.Split(Conf.HostCassaandra, ",")
	db, errs := connection.NewCassandraConnect(hosts, Conf.UserCassaandra, Conf.PassCassaandra)

	if errs != nil {
		panic(errs)
	}

	di := di.NewUseCase(db)

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
