package client

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase/dto"
)

func (b *MqttClient) PublishMessage(c *gin.Context) {
	var dto dto.Payload

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	b.Usecase.SaveMessage(&dto)
	token := b.broker.client.Publish(b.broker.topic, 1, false, dto.Message)
	token.Wait()
	fmt.Printf("Published message to topic: %s\n", b.broker.topic)
}
