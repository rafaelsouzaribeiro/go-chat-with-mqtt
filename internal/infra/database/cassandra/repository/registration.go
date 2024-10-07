package repository

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
)

func (i *CassandraRepository) Registration(user entity.User) (*entity.User, error) {

	u := i.GetUsername(user.Username)

	if u > 0 {
		return nil, fmt.Errorf("username already exists")
	}

	pg := i.GetPaginationUser()
	defer pg.Iter.Close()

	q := fmt.Sprintf(`INSERT INTO %s.users_login (id,username,password,photo,times)
						  VALUES (?, ?, ?, ?, ?)`, entity.KeySpace)

	err := i.gocql.Query(q, uuid.NewString(), user.Username, user.Password, user.Photo, time.Now()).Exec()

	if err != nil {
		return nil, err
	}

	q = fmt.Sprintf(`INSERT INTO %s.users (id,pages,username,password,photo,times)
	VALUES (?, ?, ?, ?, ?,?)`, entity.KeySpace)

	err = i.gocql.Query(q, uuid.NewString(), pg.Page, user.Username, user.Password, user.Photo, time.Now()).Exec()

	if err != nil {
		return nil, err
	}

	if pg.Iter.NumRows() == 0 {
		query := fmt.Sprintf(`INSERT INTO %s.pagination_users (id,page,total) VALUES (?,?,?)`,
			entity.KeySpace)

		err = i.gocql.Query(query, uuid.NewString(), 1, 1).Exec()

		if err != nil {
			return nil, err
		}
	} else {
		query := fmt.Sprintf(`UPDATE %s.pagination_users SET page = ?, total = ? 
							  WHERE id = ?`, entity.KeySpace)

		err = i.gocql.Query(query, pg.Page, pg.Total, pg.Id).Exec()

		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}
