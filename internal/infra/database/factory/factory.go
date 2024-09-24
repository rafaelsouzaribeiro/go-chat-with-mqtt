package factory

import "fmt"

func (f *Factory) GetConnection() (*Iconnection, error) {

	switch f.types {
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
