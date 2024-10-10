package usecase

import (
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase/dto"
)

func (i *UseCaseMessageUser) ListMessageIndex(id, receive string) (*[]dto.PayloadMesage, error) {
	list, err := i.Irepository.ListMessageIndex(id, receive)

	if err != nil {
		return nil, err
	}

	var lists []dto.PayloadMesage

	for _, v := range *list {
		lists = append(lists, dto.PayloadMesage{
			Message:  v.Message,
			Username: v.Username,
			UserId:   v.UserId,
			Times:    v.Times,
			Pages:    v.Pages,
			Receive:  v.Receive,
			Types:    v.Types,
		})

	}

	return &lists, nil
}