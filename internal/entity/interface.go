package entity

type Irepository interface {
	SaveMessage(msg *Message) error
}
