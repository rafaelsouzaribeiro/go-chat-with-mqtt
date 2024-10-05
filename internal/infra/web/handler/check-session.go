package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (o *ChatHandler) CheckSession(c *gin.Context) {
	session, err := store.Get(c.Request, "go-chat")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting session"})
		return
	}

	username, ok := session.Values["username"].(string)
	if !ok || username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User is not logged in"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"username": username})
}
