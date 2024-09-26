package client

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
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
	broker *Broker
}

func NewClient(broker *Broker) *MqttClient {
	return &MqttClient{
		broker: broker,
	}
}
