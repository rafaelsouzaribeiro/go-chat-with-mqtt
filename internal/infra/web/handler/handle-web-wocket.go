package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase/dto"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (d *ChatHandler) HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	mu.Lock()
	clients[conn] = true
	mu.Unlock()

	var currentUser dto.PayloadUser
	for {
		var msgs dto.PayloadUser
		err := conn.ReadJSON(&msgs)
		if err != nil {
			break
		}

		currentUser = msgs
		d.broadcast(msgs)
	}

	mu.Lock()
	m := &dto.PayloadUser{
		Username: currentUser.Username,
		Photo:    currentUser.Photo,
		Status:   "offline",
		Times:    currentUser.Times,
		Id:       currentUser.Id}
	d.chatUseCase.SendStatus(m)

	for conn := range clients {
		if err := conn.WriteJSON(m); err != nil {
			fmt.Println("Error sending message:", err)
			conn.Close()
			delete(clients, conn)
		}
	}

	delete(clients, conn)
	mu.Unlock()
}
