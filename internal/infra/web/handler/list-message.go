package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (o *ChatHandler) ListMessage(c *gin.Context) {
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not logged in"})
		return
	}

	id := c.Param("id")
	receive := c.Param("receive")

	output, err := o.chatUseCase.ListMessage(id, receive)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, output)
}
