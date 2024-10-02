package webserver

import "github.com/gin-gonic/gin"

type WebServer struct {
	router        *gin.Engine
	webServerPort string
}

func NewWebServer(port string) *WebServer {
	webserver := &WebServer{
		router:        gin.Default(),
		webServerPort: port,
	}

	return webserver
}
