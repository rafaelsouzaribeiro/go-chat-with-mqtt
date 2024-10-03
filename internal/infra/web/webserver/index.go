package webserver

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/infra/web/handler"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase"
)

func (w *WebServer) AddHandlerChat(order *usecase.UseCaseMessageUser) {
	chatHandler := handler.NewChatHandler(order)
	w.router.Static("/static", "../web/static")
	w.router.LoadHTMLGlob("../web/templates/*")
	w.router.GET("/list-message/:id", chatHandler.ListMessage)
	w.router.GET("/list-users/", chatHandler.Lists)
	w.router.GET("/", chatHandler.StartTemplates)
	w.router.GET("/message", chatHandler.MessageTemplates)
}

func (w *WebServer) Start() {

	w.router.Run(fmt.Sprintf(":%s", w.webServerPort))
}
