package client

import (
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase/dto"
)

func (b *MqttClient) PublishMessage(dto *dto.Payload) error {

	_, err := b.Usecase.SaveMessage(dto)

	if err != nil {
		return err
	}

	token := b.broker.Client.Publish(b.broker.Topic, 1, false, dto.Message)
	token.Wait()

	return nil
}
