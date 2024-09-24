package factory

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/configs"
)

type Factory struct {
	Types string
	Conf  *configs.Conf
}

func NewFactory(f *Factory) (*Iconnection, error) {

	switch f.Types {
	case "cassandra":
		con, err := f.GetConCassandra()

		if err != nil {
			panic(err)
		}

		c := &Iconnection{Gocql: con}

		return c, nil
	}

	return nil, fmt.Errorf("invalid driver")
}
