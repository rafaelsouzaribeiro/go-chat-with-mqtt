package usecase

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase/dto"
)

func (r *UseCaseMessageUser) CheckUser(password, username string) (*dto.PayloadUser, error) {

	user, err := r.Irepository.CheckUser(password, username)

	if err != nil {
		return nil, err
	}

	for user != nil {
		return &dto.PayloadUser{
			Username: user.Username,
			Id:       user.Id,
			Photo:    user.Photo,
			Times:    user.Times,
		}, nil
	}

	return nil, fmt.Errorf("user not found")
}
