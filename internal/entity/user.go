package entity

import (
	"github.com/gorilla/websocket"
)

type User struct {
	username string
	id       string
	conn     *websocket.Conn
}
