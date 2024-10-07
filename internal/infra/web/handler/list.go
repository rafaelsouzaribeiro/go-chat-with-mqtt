package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (o *ChatHandler) ListMessage(c *gin.Context) {
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
