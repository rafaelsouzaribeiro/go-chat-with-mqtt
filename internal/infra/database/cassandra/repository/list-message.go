package repository

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
)

func (r *CassandraRepository) ListMessage(id, receive string) (*[]entity.Message, error) {

	pg := r.GetPaginationMessageIndex(id, receive)
	defer pg.Iter.Close()
	p := pg.Page

	s := fmt.Sprintf(`SELECT id,message,pages,username,userid,times,receive,types FROM %s.messages 
	WHERE pages=? AND userid=? AND receive=? ORDER BY times ASC;`, entity.KeySpace)
	query := r.gocql.Query(s, p, id, receive)
	iter := query.Iter()
	defer iter.Close()

	var message entity.Message
	messages := make(map[string]entity.Message)
	var m []entity.Message

	for iter.Scan(&message.Id, &message.Message, &message.Pages, &message.Username,
		&message.UserId, &message.Times, &message.Receive, &message.Types) {
		message.Types = "received"
		messages[message.Id] = message
	}

	query2 := r.gocql.Query(s, p, receive, id)
	iter2 := query2.Iter()
	defer iter2.Close()

	for iter2.Scan(&message.Id, &message.Message, &message.Pages, &message.Username,
		&message.UserId, &message.Times, &message.Receive, &message.Types) {
		message.Types = "sent"
		messages[message.Id] = message
	}

	for _, msg := range messages {
		msg.PageTotal = int64(p)
		m = append(m, msg)
	}

	return &m, nil
}
