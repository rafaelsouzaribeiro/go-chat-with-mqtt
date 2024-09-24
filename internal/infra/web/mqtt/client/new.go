package client

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase"
)

type Broker struct {
	Broker   string
	Port     int
	Client   mqtt.Client
	Topic    string
	Username string
	Password string
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
