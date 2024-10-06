package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
)

func (o *ChatHandler) ActionRegistration(c *gin.Context) {
	var loginReq entity.User

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	fmt.Printf("%v", loginReq)
	c.JSON(http.StatusOK, loginReq)
}
