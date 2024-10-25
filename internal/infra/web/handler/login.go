package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (o *ChatHandler) LoginTemplates(c *gin.Context) {

	o.ClearSession(c, "go-chat")
	c.HTML(http.StatusOK, "login.html", nil)
}
