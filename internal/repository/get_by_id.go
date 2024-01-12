package repository

import (
	"database/sql"
	"fmt"
)

func (r Repository) GetById(table string, id int) (*sql.Row, error) {
	return r.SingleQuery(fmt.Sprintf("SELECT * FROM %s WHERE id = $1", table), id)
}
