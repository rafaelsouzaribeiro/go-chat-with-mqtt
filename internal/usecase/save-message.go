package usecase

import (
	"time"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase/dto"
)

func (i *UseCaseMessageUser) SaveMessage(input *dto.Payload) (*dto.Payload, error) {
	err := i.Irepository.PublishMessage(&entity.Message{
		Message:   input.Message,
		Topic:     input.Topic,
		Time:      time.Now(),
		MessageId: input.MessageId,
	})

	if err != nil {
		return nil, err
	}

	return input, nil
}
