package repository

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"

	"github.com/gocql/gocql"
)

func (i *CassandraRepository) Registration(user entity.User) (*entity.User, error) {

	u := i.GetUsername(user.Username)

	if u > 0 {
		return nil, fmt.Errorf("username already exists")
	}

	pg := i.GetPaginationUser()
	defer pg.Iter.Close()

	batch := i.gocql.NewBatch(gocql.LoggedBatch)
	id := uuid.NewString()

	q := fmt.Sprintf(`INSERT INTO %s.users_login (id,username,password,photo,times)
						  VALUES (?, ?, ?, ?, ?)`, entity.KeySpace)

	batch.Query(q, id, user.Username, user.Password, user.Photo, time.Now())

	q = fmt.Sprintf(`INSERT INTO %s.users (id,pages,username,password,photo,times)
	VALUES (?, ?, ?, ?, ?,?)`, entity.KeySpace)

	batch.Query(q, id, pg.Page, user.Username, user.Password, user.Photo, time.Now())

	if pg.Iter.NumRows() == 0 {
		query := fmt.Sprintf(`INSERT INTO %s.pagination_users (id,page,total) VALUES (?,?,?)`,
			entity.KeySpace)

		batch.Query(query, uuid.NewString(), 1, 1)

	} else {
		query := fmt.Sprintf(`UPDATE %s.pagination_users SET page = ?, total = ? 
							  WHERE id = ?`, entity.KeySpace)

		batch.Query(query, pg.Page, pg.Total, pg.Id)
	}

	if err := i.gocql.ExecuteBatch(batch); err != nil {
		return nil, err
	}

	return &user, nil
}
