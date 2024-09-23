package client

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase/dto"
)

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func (b *Broker) SetClient(pay dto.Payload, canalChan chan<- dto.Payload) {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", b.broker, b.port))
	opts.SetUsername(pay.Username)
	opts.SetPassword(pay.Password)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	token := client.Subscribe(pay.Topic, 1, func(c mqtt.Client, m mqtt.Message) {
		canalChan <- dto.Payload{
			Topic:     m.Topic(),
			Message:   string(m.Payload()),
			MessageId: m.MessageID(),
		}
	})

	b.client = client
	b.topic = pay.Topic

	token.Wait()
	fmt.Printf("Subscribed to topic: %s", pay.Topic)

}
