package repository

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
)

func (r *CassandraRepository) GetStatusUser() ([]entity.User, error) {

	s := fmt.Sprintf(`SELECT id,status,times FROM %s.users_status`, entity.KeySpace)
	query := r.gocql.Query(s)
	iter := query.Iter()
	defer iter.Close()

	var user entity.User
	var users []entity.User
	if iter.Scan(&user.Id, &user.Status, &user.Times) {
		users = append(users, user)
	}

	return users, nil
}
