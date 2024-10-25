package handler

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase/dto"
)

func (d *ChatHandler) broadcast(msg dto.PayloadUser) {
	mu.Lock()
	defer mu.Unlock()

	for conn := range clients {
		if err := conn.WriteJSON(msg); err != nil {
			fmt.Println("Erro ao enviar mensagem:", err)
			conn.Close()
			delete(clients, conn) // Remove conexão se falhar
		}
	}
}