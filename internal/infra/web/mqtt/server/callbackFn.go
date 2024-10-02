package server

import (
	"encoding/json"
	"fmt"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase/dto"
)

func (b *Broker) callbackFn(cl *mqtt.Client, sub packets.Subscription, pk packets.Packet) {
	var Payload dto.PayloadMesage

	err := json.Unmarshal(pk.Payload, &Payload)
	if err != nil {
		fmt.Printf("Failed to deserialize message: %v\n", err)
		return
	}

	dto := dto.PayloadMesage{
		Message:  Payload.Message,
		Username: Payload.Username,
		UserId:   Payload.UserId,
	}

	_, err = b.Usecase.SaveMessage(&dto)

	if err != nil {
		fmt.Printf("Failed to save message: %v\n", err)
		return
	}
}
