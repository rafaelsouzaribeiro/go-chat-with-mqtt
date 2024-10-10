package repository

import (
	"fmt"
	"sort"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
)

func (r *CassandraRepository) ListMessageIndex(id, receive string) (*[]entity.Message, error) {

	s := fmt.Sprintf(`SELECT id,message,pages,username,userid,times,receive,types FROM %s.messages 
	WHERE pages=? AND userid=? AND receive=? ORDER BY times DESC;`, entity.KeySpace)

	query := r.gocql.Query(s, entity.IndexM, id, receive)

	iter := query.Iter()
	defer iter.Close()

	var message entity.Message
	var messages []entity.Message

	for iter.Scan(&message.Id, &message.Message, &message.Pages, &message.Username,
		&message.UserId, &message.Times, &message.Receive, &message.Types) {
		message.Types = "received"

		messages = append(messages, message)
	}

	query2 := r.gocql.Query(s, entity.IndexM, receive, id)
	iter2 := query2.Iter()
	defer iter2.Close()

	for iter2.Scan(&message.Id, &message.Message, &message.Pages, &message.Username,
		&message.UserId, &message.Times, &message.Receive, &message.Types) {
		message.Types = "sent"

		messages = append(messages, message)
	}

	sort.Slice(messages, func(i, j int) bool {
		return messages[i].Times.Before(messages[j].Times)
	})

	return &messages, nil
}
