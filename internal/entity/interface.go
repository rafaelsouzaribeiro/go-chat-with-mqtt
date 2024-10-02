package entity

import "sync"

type Irepository interface {
	PublishMessage(msg *Message) error
	ListUser(id int64) (*[]User, error)
}

var (
	StartMIndex = int64(1)
	StartUIndex = int64(1)
	PerPage     = int64(20)
	KeySpace    = "chatmqtt"
	IndexU      int64
	Once        sync.Once
	IdUser      int64
)
