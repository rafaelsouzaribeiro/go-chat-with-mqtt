package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (o *ChatHandler) LoginTemplates(c *gin.Context) {

	c.HTML(http.StatusOK, "login.html", nil)
}
