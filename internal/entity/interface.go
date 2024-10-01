package entity

type Irepository interface {
	PublishMessage(msg *Message) error
	ListUser(user *User) error
}
