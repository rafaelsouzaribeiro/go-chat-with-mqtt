package repository

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
)

func (r *CassandraRepository) ListUser(id string) (*[]entity.User, error) {
	entity.Once.Do(func() { entity.IndexU = entity.StartUIndex })

	s := fmt.Sprintf(`select photo,pages,username,id,times from %s.users 
	WHERE pages=? AND id=? ORDER BY times DESC`, entity.KeySpace)
	query := r.gocql.Query(s, entity.IndexU, id)
	iter := query.Iter()
	defer iter.Close()

	var message entity.User
	var messages []entity.User

	for iter.Scan(&message.Username, &message.Photo,
		&message.Username, &message.Times) {
		messages = append(messages, message)
	}

	entity.IndexU++

	return &messages, nil
}
