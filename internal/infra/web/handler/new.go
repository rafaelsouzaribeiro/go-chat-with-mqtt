package handler

import (
	"sync"

	"github.com/gorilla/websocket"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase"
)

type ChatHandler struct {
	chatUseCase *usecase.UseCaseMessageUser
}

var (
	mu      sync.Mutex
	clients = make(map[*websocket.Conn]bool)
)

func NewChatHandler(chatUseCase *usecase.UseCaseMessageUser) *ChatHandler {
	return &ChatHandler{
		chatUseCase: chatUseCase,
	}
}
