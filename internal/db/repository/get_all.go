package repository

import (
	"fmt"

	"github.com/andresmeireles/speaker/internal/db"
	"github.com/andresmeireles/speaker/internal/db/entity"
)

func GetAll[T entity.Entity](en T) []T {
	db, err := db.GetDB()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var entities []T

	query := fmt.Sprintf("SELECT * FROM %s", en.Table())

	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		rows.Scan(entities)
	}

	return entities
}
