package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
)

var store = sessions.NewCookieStore([]byte("go-chat"))

func (o *ChatHandler) Action(c *gin.Context) {
	var loginReq entity.User

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	session, err := store.Get(c.Request, "go-chat")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting session"})
		return
	}

	passwordHash, err := o.chatUseCase.HashPassword(loginReq.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating password hash"})
		return
	}

	user, err := o.chatUseCase.Login(loginReq.Username, passwordHash)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error logging in"})
		return
	}

	session.Values["iduser"] = loginReq.Id

	err = session.Save(c.Request, c.Writer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving session"})
		return
	}

	c.JSON(http.StatusOK, user)
}
