package repository

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
)

func (r *CassandraRepository) ListUsers() (*[]entity.User, error) {
	entity.Once.Do(func() { entity.IndexU = entity.StartUIndex })

	s := fmt.Sprintf(`SELECT photo,pages,username,id,times FROM %s.users 
	WHERE pages=?;`, entity.KeySpace)

	query := r.gocql.Query(s, entity.IndexU)
	iter := query.Iter()
	defer iter.Close()

	var message entity.User
	var messages []entity.User

	for iter.Scan(&message.Photo, &message.Pages, &message.Username,
		&message.Id, &message.Times) {

		messages = append(messages, message)
	}

	entity.IndexU++

	return &messages, nil
}
