package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartTemplates() {
	router := gin.Default()
	router.LoadHTMLGlob("../web/templates/*")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{
			"title": "Login",
		})
	})
	router.Run(":8080")
}
