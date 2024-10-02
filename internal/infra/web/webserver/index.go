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
	w.router.GET("/list-user", chatHandler.List)
	w.router.GET("/", chatHandler.StartTemplates)
}

func (w *WebServer) Start() {

	w.router.Run(fmt.Sprintf(":%s", w.webServerPort))
}
