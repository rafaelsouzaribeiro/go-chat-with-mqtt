package repository

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
)

func (r *CassandraRepository) CheckUser(id string) (*entity.User, error) {

	s := fmt.Sprintf(`SELECT id,username,photo,times FROM %s.users 
					 WHERE id=?;`, entity.KeySpace)

	query := r.gocql.Query(s, id)
	iter := query.Iter()
	defer iter.Close()

	var user entity.User

	for iter.Scan(&user.Id, &user.Username,
		&user.Photo, &user.Times) {

		return &user, nil
	}

	return nil, fmt.Errorf("user not found")
}
