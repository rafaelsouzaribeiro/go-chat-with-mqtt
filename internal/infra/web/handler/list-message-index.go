package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
)

func (o *ChatHandler) ListMessageIndex(c *gin.Context) {
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

	id := c.Param("id")
	receive := c.Param("receive")

	pageStr := c.Param("page")
	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page total"})
		return
	}

	entity.IndexM = page

	output, err := o.chatUseCase.ListMessageIndex(id, receive)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, output)
}
