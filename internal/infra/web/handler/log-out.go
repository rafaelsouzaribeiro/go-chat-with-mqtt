package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (o ChatHandler) Logout(c *gin.Context) {
	o.ClearSession(c, "go-chat")
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
