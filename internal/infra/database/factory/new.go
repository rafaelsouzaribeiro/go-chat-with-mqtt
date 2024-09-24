package factory

import (
	"fmt"
)

type Factory struct {
	Factory string
}

func NewFactory(f *Factory) (*Iconnection, error) {

	switch f.Factory {
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
