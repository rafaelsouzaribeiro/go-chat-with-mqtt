package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
)

var store = sessions.NewCookieStore(securecookie.GenerateRandomKey(32))

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

	session, err := store.Get(c.Request, "go-chat")
	if err != nil {
		o.ClearSession(c, "go-chat")
	}

	session.AddFlash(user.Username)
	session.AddFlash(user.Password)

	err = session.Save(c.Request, c.Writer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error save session"})
		return
	}

	c.JSON(http.StatusOK, user)
}
