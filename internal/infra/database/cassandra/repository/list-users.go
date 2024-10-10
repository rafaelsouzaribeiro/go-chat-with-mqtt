package repository

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
)

func (r *CassandraRepository) ListUsers() (*[]entity.User, error) {
	page := r.GetPaginationUserIndex()

	s := fmt.Sprintf(`SELECT photo,pages,username,id,times FROM %s.users 
	WHERE pages=?;`, entity.KeySpace)

	query := r.gocql.Query(s, page)
	iter := query.Iter()
	defer iter.Close()

	var user entity.User
	var users []entity.User

	for iter.Scan(&user.Photo, &user.Pages, &user.Username,
		&user.Id, &user.Times) {
		user.PageTotal = page
		users = append(users, user)
	}

	return &users, nil
}
