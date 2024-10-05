package repository

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
)

func (r *CassandraRepository) Login(username string) (*entity.User, error) {
	s := fmt.Sprintf(`SELECT id,username,photo,password FROM %s.users_login
					 WHERE username=?;`, entity.KeySpace)

	query := r.gocql.Query(s, username)
	iter := query.Iter()
	defer iter.Close()

	var user entity.User

	for iter.Scan(&user.Id, &user.Username,
		&user.Photo, &user.Password) {

		return &user, nil
	}

	return nil, fmt.Errorf("user not found or incorrect credentials")
}
