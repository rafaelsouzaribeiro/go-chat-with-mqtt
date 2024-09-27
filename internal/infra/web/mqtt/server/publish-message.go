package server

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase/dto"
)

func (b *Broker) PublishMessage(c *gin.Context) {
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

	err = b.Server.Publish(b.topic, payloadJson, false, 0)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, dto)
}
