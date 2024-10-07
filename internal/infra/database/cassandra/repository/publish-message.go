package repository

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
)

func (i *CassandraRepository) PublishMessage(input *entity.Message) error {

	pg := i.GetPaginationMessage()
	defer pg.Iter.Close()

	if strings.TrimSpace(input.Message) != "" {
		q := fmt.Sprintf(`INSERT INTO %s.messages (userid,message,username,pages,times)
						  VALUES (?, ?, ?, ?, ?)`, entity.KeySpace)
		err := i.gocql.Query(q, input.UserId, input.Message,
			input.Username, pg.Page, time.Now()).Exec()

		if err != nil {
			return err
		}

		if pg.Iter.NumRows() == 0 {
			query := fmt.Sprintf(`INSERT INTO %s.pagination_messages (id,page,total) VALUES (?,?,?)`,
				entity.KeySpace)

			err = i.gocql.Query(query, uuid.NewString(), 1, 1).Exec()

			if err != nil {
				return err
			}
		} else {
			query := fmt.Sprintf(`UPDATE %s.pagination_messages SET page = ?, total = ? 
								  WHERE id = ?`, entity.KeySpace)

			err = i.gocql.Query(query, pg.Page, pg.Total, pg.Id).Exec()

			if err != nil {
				return err
			}
		}

	}

	return nil
}
