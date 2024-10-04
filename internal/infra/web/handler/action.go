package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (o *ChatHandler) Action(c *gin.Context) {
	response := gin.H{
		"message": "Action executed successfully",
		"status":  "success",
		"data": gin.H{
			"topic": "chat-topic-123",
			"user":  "John Doe",
		},
	}

	c.JSON(http.StatusOK, response)
}
