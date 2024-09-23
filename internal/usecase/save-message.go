package usecase

import (
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase/dto"
)

func (i *UseCaseMessageUser) SaveMessage(input *dto.MessageDto) (*dto.MessageDto, error) {
	err := i.Irepository.SaveMessage(&entity.Message{
		Message: input.Message,
		Type:    input.Type,
		Time:    input.Time,
	})

	if err != nil {
		return nil, err
	}

	return input, nil
}
