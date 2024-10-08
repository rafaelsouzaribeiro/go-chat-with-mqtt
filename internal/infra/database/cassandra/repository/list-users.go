package repository

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
)

func (r *CassandraRepository) ListUsers() (*[]entity.User, error) {

	p := r.GetPaginationUser()

	s := fmt.Sprintf(`SELECT photo,pages,username,id,times FROM %s.users 
	WHERE pages=?;`, entity.KeySpace)

	query := r.gocql.Query(s, p.Page)
	iter := query.Iter()
	defer iter.Close()

	var user entity.User
	var users []entity.User

	for iter.Scan(&user.Photo, &user.Pages, &user.Username,
		&user.Id, &user.Times) {

		users = append(users, user)
	}

	entity.IndexU = int64(p.Page)

	return &users, nil
}
