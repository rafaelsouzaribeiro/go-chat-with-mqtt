package factory

import (
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/configs"
)

type Factory struct {
	types string
	Conf  *configs.Conf
}

func NewFactory(types string, conf *configs.Conf) *Factory {
	return &Factory{
		types: types,
		Conf:  conf,
	}
}
