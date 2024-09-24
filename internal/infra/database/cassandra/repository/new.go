package repository

import "github.com/gocql/gocql"

type CassandraRepository struct {
	gocql *gocql.Session
}

func NewCassandraRepository(gocql *gocql.Session) *CassandraRepository {
	return &CassandraRepository{
		gocql: gocql,
	}
}
