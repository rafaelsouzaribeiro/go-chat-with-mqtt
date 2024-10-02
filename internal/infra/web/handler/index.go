package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase"
)

type ChatHandler struct {
	chatUseCase *usecase.UseCaseMessageUser
}

func NewOrderHandler(chatUseCase *usecase.UseCaseMessageUser) *ChatHandler {
	return &ChatHandler{
		chatUseCase: chatUseCase,
	}
}

func (o *ChatHandler) List(c *gin.Context) {
	output, err := o.chatUseCase.ListUser(1)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, output)
}

func StartTemplates() {
	r := gin.Default()

	r.Static("./static", "../web/static")

	r.LoadHTMLGlob("../web/templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "chat.html", nil)
	})

	r.Run(":8080")

}
