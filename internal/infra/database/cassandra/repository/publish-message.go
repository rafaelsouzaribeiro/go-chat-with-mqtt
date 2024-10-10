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

	pg := i.GetPaginationMessage(input.UserId, input.Receive)
	defer pg.Iter.Close()

	batch := i.gocql.NewBatch(gocql.LoggedBatch)
	if strings.TrimSpace(input.Message) != "" {
		q := fmt.Sprintf(`INSERT INTO %s.messages (id,userid,message,username,pages,receive,types,times)
						  VALUES (?, ?, ?, ?, ?,?,?,?)`, entity.KeySpace)
		batch.Query(q, uuid.New().String(), input.UserId, input.Message, input.Username, pg.Page, input.Receive, "", time.Now())

		if pg.Iter.NumRows() == 0 {
			query := fmt.Sprintf(`INSERT INTO %s.pagination_messages (id,page,total) VALUES (?,?,?)`,
				entity.KeySpace)
			batch.Query(query, input.UserId, 1, 1)

		} else {
			query := fmt.Sprintf(`UPDATE %s.pagination_messages SET page = ?, total = ? 
								  WHERE id = ?`, entity.KeySpace)

			batch.Query(query, pg.Page, pg.Total, input.UserId)

		}

	}

	if err := i.gocql.ExecuteBatch(batch); err != nil {
		return err
	}

	return nil
}
