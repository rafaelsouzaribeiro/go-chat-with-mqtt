package webserver

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/infra/web/handler"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase"
	"github.com/rafaelsouzaribeiro/jwt-auth/pkg/middleware"
)

func (w *WebServer) AddHandlerChat(order *usecase.UseCaseMessageUser) {
	chatHandler := handler.NewChatHandler(order)
	w.router.Static("/static", "../web/static")
	w.router.LoadHTMLGlob("../web/templates/*")
	w.router.GET("/list-message/:id", chatHandler.ListMessage)
	w.router.GET("/list-users/", chatHandler.Lists)
	w.router.GET("/", chatHandler.LoginTemplates)
	w.router.POST("/action", chatHandler.Action)

	cre, err := middleware.NewCredential(3600, "go-index", nil)

	if err != nil {
		panic(err)
	}

	w.router.GET("/index/:token", cre.AuthMiddlewareGin(), chatHandler.IndexTemplates)
}

func (w *WebServer) Start() {

	w.router.Run(fmt.Sprintf(":%s", w.webServerPort))
}
