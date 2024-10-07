package repository

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
)

func (r *CassandraRepository) GetPaginationUser() Pagination {
	var save Pagination
	var total int = 1
	var page int = 1

	s := fmt.Sprintf(`SELECT id,page,total FROM %s.pagination_users`, entity.KeySpace)
	query := r.gocql.Query(s)
	iter := query.Iter()

	if iter.Scan(&save.Id, &save.Page, &save.Total) {
		result := save.Total % int(entity.PerPage)
		total = save.Total + 1

		if result == 0 {
			page = save.Page + 1
		} else {
			page = save.Page
		}
	}

	save.Page = page
	save.Total = total
	save.Iter = iter

	return save
}
