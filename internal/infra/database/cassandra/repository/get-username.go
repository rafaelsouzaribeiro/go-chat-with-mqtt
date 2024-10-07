package repository

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
)

func (i *CassandraRepository) GetUsername(username string) int {

	s := fmt.Sprintf(`SELECT * FROM %s.users_login WHERE username=?`, entity.KeySpace)
	query := i.gocql.Query(s, username)
	iter := query.Iter()

	return iter.NumRows()
}
