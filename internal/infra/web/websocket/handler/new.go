package handler

import "github.com/rafaelsouzaribeiro/whatsapp-clone-in-go/internal/usecase"

type MessageUserHanlder struct {
	messageUSerUsecase *usecase.UseCaseMessageUser
}

func NewMessageUserHanlder(messageUSerUsecase *usecase.UseCaseMessageUser) *MessageUserHanlder {
	return &MessageUserHanlder{
		messageUSerUsecase: messageUSerUsecase,
	}
}
