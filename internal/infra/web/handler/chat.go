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

	result := session.Flashes()

	if result == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "value not set"})
		return
	}

	user, err := o.chatUseCase.CheckUser(result[1].(string), result[0].(string))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not logged in"})
		return
	}

	data := gin.H{
		"topic":       viper.GetString("TOPIC_MQTT"),
		"idUser":      user.Id,
		"username":    user.Username,
		"photo":       user.Photo,
		"hostname":    viper.GetString("HOST_MQTT_WEBSOCKET"),
		"port":        viper.GetInt("PORT_MQTT_WEBSOCKET"),
		"usernameCon": viper.GetString("USERNAME_MQTT"),
		"password":    viper.GetString("PASSWORD_MQTT"),
	}

	c.HTML(http.StatusOK, "chat.html", data)
}
