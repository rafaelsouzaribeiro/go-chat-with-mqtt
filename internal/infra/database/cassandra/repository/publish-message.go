package repository

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
)

func (i *CassandraRepository) PublishMessage(input *entity.Message) error {

	pg := i.GetPaginationMessage()
	defer pg.Iter.Close()

	batch := i.gocql.NewBatch(gocql.LoggedBatch)

	if strings.TrimSpace(input.Message) != "" {
		q := fmt.Sprintf(`INSERT INTO %s.messages (userid,message,username,pages,receive,types,times)
						  VALUES (?, ?, ?, ?, ?,?,?)`, entity.KeySpace)
		batch.Query(q, input.UserId, input.Message, input.Username, pg.Page, input.Receive, "received", time.Now())
		batch.Query(q, input.Receive, input.Message, input.Username, pg.Page, input.UserId, "sent", time.Now())

		if pg.Iter.NumRows() == 0 {
			query := fmt.Sprintf(`INSERT INTO %s.pagination_messages (id,page,total) VALUES (?,?,?)`,
				entity.KeySpace)
			batch.Query(query, uuid.NewString(), 1, 1)

		} else {
			query := fmt.Sprintf(`UPDATE %s.pagination_messages SET page = ?, total = ? 
								  WHERE id = ?`, entity.KeySpace)

			batch.Query(query, pg.Page, pg.Total, pg.Id)

		}

	}

	if err := i.gocql.ExecuteBatch(batch); err != nil {
		return err
	}

	return nil
}
