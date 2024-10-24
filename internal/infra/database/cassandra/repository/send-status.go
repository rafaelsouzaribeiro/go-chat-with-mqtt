package repository

import (
	"fmt"
	"time"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
)

func (i *CassandraRepository) SendStatus(id, status string) {

	s := fmt.Sprintf(`SELECT * FROM %s.users_status WHERE id=?`, entity.KeySpace)
	query := i.gocql.Query(s, id)
	iter := query.Iter()
	defer iter.Close()

	if iter.NumRows() == 0 {
		q := fmt.Sprintf(`INSERT INTO %s.users_status (id,status,times)
		VALUES (?, ? , ?)`, entity.KeySpace)

		_ = i.gocql.Query(q, id, "online", time.Now()).Exec()

	}

	if iter.NumRows() == 1 {
		q := fmt.Sprintf(`UPDATE %s.users_status SET  status=?,times=? WHERE id=?`, entity.KeySpace)

		_ = i.gocql.Query(q, status, time.Now(), id).Exec()

	}

}
