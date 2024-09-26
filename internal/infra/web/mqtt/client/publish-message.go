package client

import (
	"encoding/json"
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

	payloadJson, err := json.Marshal(dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize message"})
		return
	}

	token := b.broker.Client.Publish(b.broker.Topic, 1, false, payloadJson)
	token.Wait()

	c.JSON(http.StatusCreated, dto)
}
