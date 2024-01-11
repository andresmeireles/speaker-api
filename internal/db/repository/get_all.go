package repository

import (
	"database/sql"
	"fmt"
)

func (r Repository[T]) GetAll() (*sql.Rows, error) {
	en := *new(T)

	return r.Query(fmt.Sprintf("SELECT * FROM %s", en.Table()))
}
