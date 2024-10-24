package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
)

func (o *ChatHandler) LisUsersStatus(c *gin.Context) {
	session, err := store.Get(c.Request, "go-chat")

	if err != nil {
		o.ClearSession(c, "go-chat")
	}

	result := session.Flashes()

	if result == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "value not set"})
		return
	}

	_, err = o.chatUseCase.CheckUser(result[1].(string), result[0].(string))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not logged in"})
		return
	}

	output, err := o.chatUseCase.GetStatusUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	var out []entity.User
	for _, v := range *output {
		out = append(out, entity.User{
			Id:     v.Id,
			Times:  v.Times,
			Status: v.Status,
		})
	}

	c.JSON(http.StatusOK, out)
}
