package client

import (
	"encoding/json"
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

func (b *MqttClient) Connect(canalChan chan<- dto.Payload) {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", b.broker.Broker, b.broker.Port))
	opts.SetUsername(b.broker.Username)
	opts.SetPassword(b.broker.Password)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	token := client.Subscribe(b.broker.Topic, 1, func(c mqtt.Client, m mqtt.Message) {
		var payload dto.Payload

		err := json.Unmarshal(m.Payload(), &payload)
		if err != nil {
			fmt.Printf("Failed to deserialize message: %v\n", err)
			return
		}

		canalChan <- payload
	})

	b.broker.Client = client

	token.Wait()
	fmt.Printf("Subscribed to topic: %s", b.broker.Topic)

}
