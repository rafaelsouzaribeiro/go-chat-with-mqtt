package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (o *ChatHandler) StartTemplates(c *gin.Context) {
	c.HTML(http.StatusOK, "chat.html", nil)
}
