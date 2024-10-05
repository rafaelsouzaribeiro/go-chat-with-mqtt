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
		o.ClearSession(c, "go-chat")
	}

	id, ok := session.Values["idUser"]
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "value not set"})
		return
	}

	user, err := o.chatUseCase.CheckUser(id.(string))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not logged in"})
		return
	}

	data := gin.H{
		"topic":    viper.GetString("TOPIC_MQTT"),
		"idUser":   id.(string),
		"username": user.Username,
		"photo":    user.Photo,
	}

	c.HTML(http.StatusOK, "chat.html", data)
}
