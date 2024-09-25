package repository

import (
	"github.com/gocql/gocql"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/infra/database/factory"
)

type CassandraRepository struct {
	gocql *gocql.Session
}

func NewCassandraRepository(db *factory.Iconnection) *CassandraRepository {
	return &CassandraRepository{
		gocql: db.Gocql,
	}
}