package repository

import (
	"database/sql"
	"fmt"

	"github.com/andresmeireles/speaker/internal/db/entity"
)

func GetAll[T entity.Entity]() (*sql.Rows, error) {
	en := *new(T)

	return Query(fmt.Sprintf("SELECT * FROM %s", en.Table()))
}
