package client

import mqtt "github.com/eclipse/paho.mqtt.golang"

type Broker struct {
	broker   string
	port     int
	client   mqtt.Client
	topic    string
	username string
	password string
}

func NewBroker(broker, topic, username, password string, client mqtt.Client, port int) *Broker {
	return &Broker{
		broker:   broker,
		port:     port,
		topic:    topic,
		username: username,
		password: password,
		client:   client,
	}
}
