package handler

import (
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase"
)

type ChatHandler struct {
	chatUseCase *usecase.UseCaseMessageUser
}

func NewChatHandler(chatUseCase *usecase.UseCaseMessageUser) *ChatHandler {
	return &ChatHandler{
		chatUseCase: chatUseCase,
	}
}
