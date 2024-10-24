package repository

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
)

func (i *CassandraRepository) VerifyStatus(id string) error {

	q := fmt.Sprintf(`SELECT * FROM %s.users_status WHERE id=?`, entity.KeySpace)

	query := i.gocql.Query(q, id)
	iter := query.Iter()

	var user entity.User
	if iter.Scan(&user.Id, &user.Status) {

	}

	return nil
}
