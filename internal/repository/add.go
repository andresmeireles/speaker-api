package repository

import (
	"fmt"

	"github.com/andresmeireles/speaker/internal/db"
)

func (r Repository) Add(en db.Entity) error {
	db, err := r.conn.GetDB()
	if err != nil {
		return err
	}

	defer db.Close()

	keys, interrogations, values := r.split(en)
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
