package usecase

import "github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase/dto"

func (i *UseCaseMessageUser) Login(username, password string) (*dto.PayloadUser, error) {
	user, err := i.Irepository.Login(username, password)

	if err != nil {
		return nil, err
	}

	return &dto.PayloadUser{
		Username: user.Photo,
		Id:       user.Id,
		Photo:    user.Photo,
		Times:    user.Times,
	}, nil
}
