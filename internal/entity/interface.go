package entity

import "sync"

type Irepository interface {
	PublishMessage(msg *Message) error
	ListMessage(id string) (*[]Message, error)
	ListUsers() (*[]User, error)
	Login(username string) (*User, error)
	CheckUser(id string) (*User, error)
}

var (
	StartUIndex = int64(1)
	PerPage     = int64(20)
	KeySpace    = "chatmqtt"
	IndexU      = int64(1)
	IndexM      = int64(1)
	Once        sync.Once
	IdUser      int64
)
