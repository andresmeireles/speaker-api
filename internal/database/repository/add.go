package repository

import (
	"fmt"

	"github.com/andresmeireles/speaker/internal/database"
	"github.com/andresmeireles/speaker/internal/database/entity"
)

func Add[T entity.Entity](en T) error {
	db, err := database.GetDB()

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
