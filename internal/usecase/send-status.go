package usecase

import (
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase/dto"
)

func (i *UseCaseMessageUser) SendStatus(user dto.PayloadUser) {

	i.Irepository.SendStatus(entity.User{
		Username: user.Username,
		Id:       user.Id,
		Photo:    user.Photo,
		Times:    user.Times,
		Status:   user.Status,
	})

}
