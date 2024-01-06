package repository

import (
	"database/sql"

	"github.com/andresmeireles/speaker/internal/db"
)

func SingleQuery(q string, args any) (*sql.Row, error) {
	db, err := db.GetDB()
	defer db.Close()

	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare(q)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(args)
	if row.Err() != nil {
		return nil, row.Err()
	}

	return row, nil
}

func Query(q string, args ...any) (*sql.Rows, error) {
	db, err := db.GetDB()

	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare(q)

	if err != nil {
		return nil, err
	}

	return stmt.Query(args...)
}
