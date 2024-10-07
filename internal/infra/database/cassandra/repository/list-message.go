package repository

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
)

func (r *CassandraRepository) ListMessage(id, receive string) (*[]entity.Message, error) {

	s := fmt.Sprintf(`SELECT message,pages,username,userid,times,receive FROM %s.messages 
	WHERE pages=? AND userid=? AND receive=? ORDER BY times ASC;`, entity.KeySpace)

	query := r.gocql.Query(s, entity.IndexM, id, receive)
	iter := query.Iter()
	defer iter.Close()

	var message entity.Message
	var messages []entity.Message

	for iter.Scan(&message.Message, &message.Pages, &message.Username,
		&message.UserId, &message.Times, &message.Receive) {

		messages = append(messages, message)
	}

	return &messages, nil
}
