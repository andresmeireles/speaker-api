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

	_, err = db.Exec(query, values...)

	if err != nil {
		return err
	}

	return nil
}
