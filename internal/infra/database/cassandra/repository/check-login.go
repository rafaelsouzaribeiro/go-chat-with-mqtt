package repository

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
)

func (r *CassandraRepository) CheckUser(password, username string) (*entity.User, error) {

	s := fmt.Sprintf(`SELECT id,username,photo,times FROM %s.users_login 
					 WHERE password=? AND username=?;`, entity.KeySpace)

	query := r.gocql.Query(s, password, username)
	iter := query.Iter()
	defer iter.Close()

	var user entity.User

	for iter.Scan(&user.Id, &user.Username,
		&user.Photo, &user.Times) {

		return &user, nil
	}

	return nil, fmt.Errorf("user not found")
}
