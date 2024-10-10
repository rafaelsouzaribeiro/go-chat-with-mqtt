package repository

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
)

func (r *CassandraRepository) ListMessage(id, receive string) (*[]entity.Message, error) {

	pg := r.GetPaginationMessage(id, receive)
	defer pg.Iter.Close()
	p := pg.Page

	s := fmt.Sprintf(`SELECT message,pages,username,userid,times,receive,types FROM %s.messages 
	WHERE pages=? AND userid=? AND receive=? ORDER BY times ASC;`, entity.KeySpace)
	query := r.gocql.Query(s, p, id, receive)
	iter := query.Iter()
	defer iter.Close()

	var message entity.Message
	var messages []entity.Message

	for iter.Scan(&message.Message, &message.Pages, &message.Username,
		&message.UserId, &message.Times, &message.Receive, &message.Types) {
		message.Types = "received"

		messages = append(messages, message)
	}

	query2 := r.gocql.Query(s, p, receive, id)
	iter2 := query2.Iter()
	defer iter2.Close()

	for iter2.Scan(&message.Message, &message.Pages, &message.Username,
		&message.UserId, &message.Times, &message.Receive, &message.Types) {
		message.Types = "sent"

		messages = append(messages, message)
	}

	return &messages, nil
}
