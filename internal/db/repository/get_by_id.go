package repository

import (
	"database/sql"
	"fmt"

	"github.com/andresmeireles/speaker/internal/db"
	"github.com/andresmeireles/speaker/internal/db/entity"
)

func GetById[T entity.Entity](id int, en T) (*T, error) {
	var e T

	db, err := db.GetDB()

	defer db.Close()

	if err != nil {
		panic(err)
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", en.Table())

	row := db.QueryRow(query, id)

	if err := row.Scan(e); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("entity with id %d not found", id)
		}
	}

	return &e, nil
}
