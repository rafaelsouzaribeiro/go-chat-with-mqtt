package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase/dto"
)

func (o *ChatHandler) ActionRegistration(c *gin.Context) {
	var loginReq entity.User

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	println(loginReq.Username, loginReq.Photo, loginReq.Password)
	user, err := o.chatUseCase.Registration(&dto.PayloadUser{
		Username: loginReq.Username,
		Photo:    loginReq.Photo,
		Password: loginReq.Password,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
