package client

import (
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

	_, err := b.Usecase.SaveMessage(&dto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	token := b.broker.Client.Publish(b.broker.Topic, 1, false, dto.Message)
	token.Wait()

	c.JSON(http.StatusCreated, dto)
}
