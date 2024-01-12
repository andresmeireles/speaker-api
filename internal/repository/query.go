package repository

import (
	"database/sql"
)

func (r Repository[T]) SingleQuery(q string, args ...any) (*sql.Row, error) {
	db, err := r.conn.GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(args...)
	if row.Err() != nil {
		return nil, row.Err()
	}

	return row, nil
}

func (r Repository[T]) Query(q string, args ...any) (*sql.Rows, error) {
	db, err := r.conn.GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return nil, err
	}

	return stmt.Query(args...)
}
