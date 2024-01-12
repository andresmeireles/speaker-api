package repository

import (
	"database/sql"
	"fmt"
)

func (r Repository) GetAll(table string) (*sql.Rows, error) {
	return r.Query(fmt.Sprintf("SELECT * FROM %s", table))
}
