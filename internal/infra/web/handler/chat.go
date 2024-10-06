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

	username, ok := session.Values["username"]
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "value not set"})
		return
	}

	password, ok := session.Values["password"]
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "value not set"})
		return
	}

	user, err := o.chatUseCase.CheckUser(password.(string), username.(string))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not logged in"})
		return
	}

	data := gin.H{
		"topic":    viper.GetString("TOPIC_MQTT"),
		"idUser":   user.Id,
		"username": user.Username,
		"photo":    user.Photo,
	}

	c.HTML(http.StatusOK, "chat.html", data)
}
