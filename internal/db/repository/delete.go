package repository

import (
	"fmt"

	"github.com/andresmeireles/speaker/internal/db"
	"github.com/andresmeireles/speaker/internal/db/entity"
)

func Delete(en entity.Entity) error {
	db, err := db.GetDB()

	if err != nil {
		return err
	}

	defer db.Close()

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", en.Table())

	_, err = db.Exec(query, en.GetId())

	return err
}
