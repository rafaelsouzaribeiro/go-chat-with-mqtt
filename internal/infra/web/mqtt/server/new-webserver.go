package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type WebServer struct {
	Router *gin.Engine
	Port   string
}

func NewWebServer(port string) *WebServer {
	webserver := &WebServer{
		Router: gin.Default(),
		Port:   port,
	}

	return webserver
}

func (w *WebServer) Start() {
	err := w.Router.Run(fmt.Sprintf(":%s", w.Port))

	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
