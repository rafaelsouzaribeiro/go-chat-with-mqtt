package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (o ChatHandler) ClearSession(c *gin.Context, session string) {
	cookie := &http.Cookie{
		Name:   session,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(c.Writer, cookie)
}
