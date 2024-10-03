package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (o *ChatHandler) Lists(c *gin.Context) {

	output, err := o.chatUseCase.ListUsers()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, output)
}
