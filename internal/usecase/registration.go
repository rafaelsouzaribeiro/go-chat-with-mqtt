package usecase

import (
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase/dto"
)

func (i *UseCaseMessageUser) Registration(input *dto.PayloadUser) (*dto.PayloadUser, error) {
	user, err := i.Irepository.Registration(entity.User{
		Username: input.Username,
		Photo:    input.Photo,
		Password: input.Password,
	})

	if err != nil {
		return nil, err
	}

	return &dto.PayloadUser{
		Username: user.Username,
		Photo:    user.Photo,
	}, nil
}
