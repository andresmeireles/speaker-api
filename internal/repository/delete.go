package repository

import (
	"fmt"

	"github.com/andresmeireles/speaker/internal/db"
)

func (r Repository) Delete(en db.Entity) error {
	db, err := r.conn.GetDB()
	if err != nil {
		return err
	}
	defer db.Close()

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", en.Table())
	_, err = db.Exec(query, en.GetId())

	return err
}
