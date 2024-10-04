package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func (o *ChatHandler) IndexTemplates(c *gin.Context) {
	viper.AutomaticEnv()

	data := gin.H{
		"topic": viper.GetString("TOPIC_MQTT"),
	}

	c.HTML(http.StatusOK, "chat.html", data)
}
