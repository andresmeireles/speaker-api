package repository

import (
	"database/sql"
	"fmt"

	"github.com/andresmeireles/speaker/internal/db/entity"
)

func GetById[T entity.Entity](id int) (*sql.Row, error) {
	var e T

	return SingleQuery(fmt.Sprintf("SELECT * FROM %s WHERE id = $1", e.Table()), id)
}
