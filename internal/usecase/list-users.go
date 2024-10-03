package usecase

import (
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase/dto"
)

func (i *UseCaseMessageUser) ListUsers() (*[]dto.PayloadUser, error) {
	list, err := i.Irepository.ListUsers()

	if err != nil {
		return nil, err
	}

	var lists []dto.PayloadUser

	for _, v := range *list {
		lists = append(lists, dto.PayloadUser{
			Username: v.Username,
			Id:       v.Id,
			Photo:    v.Photo,
			Time:     v.Times,
		})

	}

	return &lists, nil
}
