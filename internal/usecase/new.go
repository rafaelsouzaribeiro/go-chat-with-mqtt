package usecase

import "github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"

type UseCaseMessageUser struct {
	Irepository entity.Irepository
}

func NewUsecase(Irepository entity.Irepository) *UseCaseMessageUser {
	return &UseCaseMessageUser{
		Irepository: Irepository,
	}
}
