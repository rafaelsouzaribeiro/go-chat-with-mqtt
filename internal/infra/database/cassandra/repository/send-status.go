package repository

import (
	"fmt"
	"time"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
)

func (i *CassandraRepository) SendStatus(user entity.User) {

	s := fmt.Sprintf(`SELECT * FROM %s.users_status WHERE id=?`, entity.KeySpace)
	query := i.gocql.Query(s, user.Id)
	iter := query.Iter()
	defer iter.Close()

	if iter.NumRows() == 0 {
		q := fmt.Sprintf(`INSERT INTO %s.users_status (id,status,times,photo,username)
		VALUES (?, ? , ?,?,?)`, entity.KeySpace)

		_ = i.gocql.Query(q, user.Id, "online", time.Now(), user.Photo, user.Username).Exec()

	}

	if iter.NumRows() == 1 {
		q := fmt.Sprintf(`UPDATE %s.users_status SET  
							status=?,times=?,photo=?,username=?
							WHERE id=?`, entity.KeySpace)

		_ = i.gocql.Query(q, user.Status, time.Now(), user.Photo, user.Username, user.Id).Exec()

	}

}
