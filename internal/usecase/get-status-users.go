package usecase

import (
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase/dto"
)

func (r *UseCaseMessageUser) GetStatusUser() ([]dto.PayloadUser, error) {

	query, err := r.Irepository.GetStatusUser()

	if err != nil {
		return nil, err
	}

	var users []dto.PayloadUser

	for _, obj := range query {
		users = append(users, dto.PayloadUser{
			Id:     obj.Id,
			Times:  obj.Times,
			Status: obj.Status,
		})
	}

	return users, nil
}
