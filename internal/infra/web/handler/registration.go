package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (o *ChatHandler) RegistrationTemplates(c *gin.Context) {

	c.HTML(http.StatusOK, "registration.html", nil)
}
