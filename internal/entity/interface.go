package entity

const PerPage = 20

type Irepository interface {
	PublishMessage(msg *Message) error
	ListUser(id int64) (*User, error)
}
