package repository

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
)

func (r *CassandraRepository) ListUsersIndex() (*[]entity.User, error) {

	entity.IndexU--

	s := fmt.Sprintf(`SELECT photo,pages,username,id,times FROM %s.users 
	WHERE pages=?;`, entity.KeySpace)
	fmt.Printf("%s,%d", s, entity.IndexU)
	query := r.gocql.Query(s, entity.IndexU)
	iter := query.Iter()
	defer iter.Close()

	var user entity.User
	var users []entity.User

	for iter.Scan(&user.Photo, &user.Pages, &user.Username,
		&user.Id, &user.Times) {

		users = append(users, user)
	}

	return &users, nil
}
