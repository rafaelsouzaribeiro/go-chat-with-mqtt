package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
)

func (o *ChatHandler) Action(c *gin.Context) {
	var loginReq entity.User

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	user, err := o.chatUseCase.Login(loginReq.Username)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error logging in"})
		return
	}

	if !o.chatUseCase.CheckPassword(user.Password, loginReq.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error logging in"})
		return
	}

	c.JSON(http.StatusOK, user)
}
