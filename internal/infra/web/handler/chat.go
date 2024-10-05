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

	a, ok := session.Values["userId"]
	if !ok {
		http.Error(c.Writer, "value not set", http.StatusInternalServerError)
		return
	}

	data := gin.H{
		"topic":  viper.GetString("TOPIC_MQTT"),
		"idUser": a,
	}

	c.HTML(http.StatusOK, "chat.html", data)
}
