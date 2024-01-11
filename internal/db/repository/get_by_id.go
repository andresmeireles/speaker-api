package repository

import (
	"database/sql"
	"fmt"
)

func (r Repository[T]) GetById(id int) (*sql.Row, error) {
	var e T

	return r.SingleQuery(fmt.Sprintf("SELECT * FROM %s WHERE id = $1", e.Table()), id)
}
