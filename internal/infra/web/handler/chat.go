package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func (o *ChatHandler) IndexTemplates(c *gin.Context) {
	viper.AutomaticEnv()

	session, err := store.Get(c.Request, "go-chat")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting session"})
		return
	}

	id, ok := session.Values["iduser"].(string)
	if !ok || id == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User is not logged in"})
		return
	}

	user, err := o.chatUseCase.CheckUser(id)

	if err != nil {
		panic(err)
	}

	data := gin.H{
		"topic":    viper.GetString("TOPIC_MQTT"),
		"username": user.Username,
		"photo":    user.Photo,
		"id":       user.Id,
		"times":    user.Times,
	}

	c.HTML(http.StatusOK, "chat.html", data)
}
