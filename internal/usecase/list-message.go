package usecase

import (
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase/dto"
)

func (i *UseCaseMessageUser) ListMessage(id, receive string) (*[]dto.PayloadMesage, error) {
	list, err := i.Irepository.ListMessage(id, receive)

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
		})

	}

	return &lists, nil
}
