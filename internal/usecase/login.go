package usecase

import "github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase/dto"

func (i *UseCaseMessageUser) Login(username string) (*dto.PayloadUser, error) {
	user, err := i.Irepository.Login(username)

	if err != nil {
		return nil, err
	}

	return &dto.PayloadUser{
		Username: user.Username,
		Id:       user.Id,
		Photo:    user.Photo,
		Times:    user.Times,
		Password: user.Password,
	}, nil
}
