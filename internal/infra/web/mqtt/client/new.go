package client

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase"
)

type Broker struct {
	broker   string
	port     int
	client   mqtt.Client
	topic    string
	username string
	password string
}

type MqttClient struct {
	Usecase *usecase.UseCaseMessageUser
	broker  *Broker
}

func NewClient(usecase *usecase.UseCaseMessageUser, broker *Broker) *MqttClient {
	return &MqttClient{
		Usecase: usecase,
		broker:  broker,
	}
}
