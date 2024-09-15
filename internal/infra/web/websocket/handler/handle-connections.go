package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/rafaelsouzaribeiro/whatsapp-clone-in-go/internal/usecase/dto"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h *MessageUserHanlder) HandleConnections() {
	router := gin.Default()
	router.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer conn.Close()
		for {
			var msgs dto.MessageDto
			err := conn.ReadJSON(&msgs)
			if err != nil {
				break
			}

			conn.WriteJSON(dto.MessageDto{
				Message: "Testar",
				Type:    "message",
				Time:    time.Now(),
			})
		}
	})
	router.Run(":7007")
}
