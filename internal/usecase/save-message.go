package usecase

import (
	"time"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase/dto"
)

func (i *UseCaseMessageUser) SaveMessage(input *dto.PayloadMesage) (*dto.PayloadMesage, error) {
	err := i.Irepository.PublishMessage(&entity.Message{
		Message:  input.Message,
		Times:    time.Now(),
		Username: input.Username,
		UserId:   input.UserId,
		Receive:  input.Receive,
	})

	if err != nil {
		return nil, err
	}

	return input, nil
}
