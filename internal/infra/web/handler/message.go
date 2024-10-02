package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (o *ChatHandler) MessageTemplates(c *gin.Context) {
	c.HTML(http.StatusOK, "message.html", nil)
}
