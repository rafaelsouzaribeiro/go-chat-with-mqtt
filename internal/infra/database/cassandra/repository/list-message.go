package repository

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
)

func (r *CassandraRepository) ListMessage(id string) (*[]entity.Message, error) {

	s := fmt.Sprintf(`SELECT message,pages,username,userid,times FROM %s.messages 
	WHERE pages=? AND userid=?;`, entity.KeySpace)

	query := r.gocql.Query(s, entity.IndexM, id)
	iter := query.Iter()
	defer iter.Close()

	var message entity.Message
	var messages []entity.Message

	for iter.Scan(&message.Message, &message.Pages, &message.Username,
		&message.UserId, &message.Times) {

		messages = append(messages, message)
	}

	entity.IndexM++

	return &messages, nil
}
