package factory

import (
	"github.com/gocql/gocql"
)

const (
	Cassandra = "cassandra"
)

type Iconnection struct {
	Gocql *gocql.Session
}
