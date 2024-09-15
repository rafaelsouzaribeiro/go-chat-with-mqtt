package usecase

import (
	"github.com/rafaelsouzaribeiro/whatsapp-clone-in-go/internal/entity"
	"github.com/rafaelsouzaribeiro/whatsapp-clone-in-go/internal/usecase/dto"
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
