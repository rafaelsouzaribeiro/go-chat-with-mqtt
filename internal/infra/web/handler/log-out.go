package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (o ChatHandler) Logout(c *gin.Context) {
	session, err := store.Get(c.Request, "go-chat")
	if err != nil {
		o.ClearSession(c, "go-chat")
	}

	result := session.Flashes()

	if result == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "value not set"})
		return
	}

	user, err := o.chatUseCase.Login(result[0].(string))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error finding user"})
		return
	}

	user.Status = "offline"
	o.chatUseCase.SendStatus(*user)

	o.ClearSession(c, "go-chat")
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
