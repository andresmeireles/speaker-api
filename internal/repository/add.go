package repository

import (
	"fmt"
)

func (r Repository[T]) Add(en T) error {
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
