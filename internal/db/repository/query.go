package repository

import (
	"database/sql"

	"github.com/andresmeireles/speaker/internal/db"
)

func SingleQuery(q string, args any) *sql.Row {
	db, err := db.GetDB()
	defer db.Close()

	if err != nil {
		panic(err) // TODO: fazer algo melhor com isso
	}

	stmt, err := db.Prepare(q)

	if err != nil {
		panic(err)
	}

	return stmt.QueryRow(args)
}

func Query(q string, args ...any) (*sql.Rows, error) {
	db, err := db.GetDB()

	if err != nil {
		panic(err) // TODO: fazer algo melhor com isso
	}

	stmt, err := db.Prepare(q)

	if err != nil {
		panic(err)
	}

	return stmt.Query(args...)
}
