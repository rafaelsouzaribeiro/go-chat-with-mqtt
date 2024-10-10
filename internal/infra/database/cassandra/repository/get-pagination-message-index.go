package repository

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
)

func (r *CassandraRepository) GetPaginationMessageIndex(idUser, receive string) Pagination {
	var save Pagination

	s := fmt.Sprintf(`SELECT id,page,total FROM %s.pagination_messages WHERE id=?`, entity.KeySpace)
	query := r.gocql.Query(s, idUser)
	iter := query.Iter()
	var tempPage int

	if iter.Scan(&save.Id, &save.Page, &save.Total) {
		tempPage += save.Page
	}

	iter.Close()

	query2 := r.gocql.Query(s, receive)
	iter = query2.Iter()
	defer iter.Close()

	if iter.Scan(&save.Id, &save.Page, &save.Total) {
		tempPage += save.Page
	}

	save.Page = tempPage

	return save
}
