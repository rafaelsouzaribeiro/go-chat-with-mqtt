package entity

type Irepository interface {
	PublishMessage(msg *Message) error
	ListMessage(id, receive string) (*[]Message, error)
	ListUsers() (*[]User, error)
	Login(username string) (*User, error)
	CheckUser(password, username string) (*User, error)
	Registration(user User) (*User, error)
	ListMessageIndex(id, receive string) (*[]Message, error)
}

var (
	PerPage  = int64(10)
	KeySpace = "chatmqtt"
	IndexU   = int64(1)
	IndexM   = int64(1)
)
