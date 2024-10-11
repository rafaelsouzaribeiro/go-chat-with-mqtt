package repository

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
)

func (r *CassandraRepository) GetPaginationMessageIndex(idUser, receive string) Pagination {
	var save Pagination
	var PageS int
	var PageR int

	s := fmt.Sprintf(`SELECT id,page,total FROM %s.pagination_messages WHERE id=?`, entity.KeySpace)
	query1 := r.gocql.Query(s, fmt.Sprintf("%s|%s", idUser, receive))
	iter1 := query1.Iter()

	if iter1.Scan(&save.Id, &save.Page, &save.Total) {
		PageS = save.Page
	}

	iter1.Close()
	query2 := r.gocql.Query(s, fmt.Sprintf("%s|%s", receive, idUser))
	iter2 := query2.Iter()

	if iter2.Scan(&save.Id, &save.Page, &save.Total) {
		PageR = save.Page
	}

	iter2.Close()

	if PageS >= PageR {
		save.Page = PageS
	}

	if PageS <= PageR {
		save.Page = PageR
	}

	return save
}
