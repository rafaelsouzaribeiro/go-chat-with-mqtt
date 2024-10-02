package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (o *ChatHandler) List(c *gin.Context) {
	id := c.Param("id")

	idstr, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	output, err := o.chatUseCase.ListUser(idstr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, output)
}
