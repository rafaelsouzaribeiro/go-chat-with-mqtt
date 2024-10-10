package repository

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
)

func (r *CassandraRepository) GetPaginationUserIndex() int64 {
	var save Pagination

	s := fmt.Sprintf(`SELECT id,page,total FROM %s.pagination_users`, entity.KeySpace)
	query := r.gocql.Query(s)
	iter := query.Iter()
	defer iter.Close()

	if iter.Scan(&save.Id, &save.Page, &save.Total) {

	}

	return int64(save.Page)
}
