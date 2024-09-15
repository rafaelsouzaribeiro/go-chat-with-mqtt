package usecase

import "github.com/rafaelsouzaribeiro/whatsapp-clone-in-go/internal/entity"

type UseCaseMessageUser struct {
	Irepository entity.Irepository
}

func NewUsecase(Irepository entity.Irepository) *UseCaseMessageUser {
	return &UseCaseMessageUser{
		Irepository: Irepository,
	}
}
