package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartTemplates() {
	r := gin.Default()

	r.Static("./static", "../web/static")

	r.LoadHTMLGlob("../web/templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "chat.html", nil)
	})

	r.Run(":8080")

}
