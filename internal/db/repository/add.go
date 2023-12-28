package repository

import (
	"fmt"

	"github.com/andresmeireles/speaker/internal/db"
	"github.com/andresmeireles/speaker/internal/db/entity"
)

func Add[T entity.Entity](en T) error {
	db, err := db.GetDB()

	if err != nil {
		return err
	}

	defer db.Close()

	keys, interrogations, values := Split(en)
	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		en.Table(),
		keys,
		interrogations,
	)
	stmt, err := db.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(values...)

	if err != nil {
		return err
	}

	return nil
}
