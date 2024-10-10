package repository

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
)

func (r *CassandraRepository) GetPaginationMessageIndex(idUser, receive string) Pagination {
	var save Pagination
	var tempPage int

	s := fmt.Sprintf(`SELECT id,page,total FROM %s.pagination_messages WHERE id=?`, entity.KeySpace)

	query1 := r.gocql.Query(s, idUser)
	iter1 := query1.Iter()

	if iter1.Scan(&save.Id, &save.Page, &save.Total) {
		tempPage += save.Page
	}

	iter1.Close()

	query2 := r.gocql.Query(s, receive)
	iter2 := query2.Iter()

	if iter2.Scan(&save.Id, &save.Page, &save.Total) {
		tempPage += save.Page
	}

	iter2.Close()

	save.Page = tempPage

	return save
}
