package repository

import (
	"database/sql"

	"github.com/andresmeireles/speaker/internal/db"
)

type RepositoryInterface interface {
	Add(en db.Entity) error
	Delete(en db.Entity) error
	GetAll(table string) (*sql.Rows, error)
	GetById(table string, id int) (*sql.Row, error)
	SingleQuery(q string, args ...any) (*sql.Row, error)
	Query(q string, args ...any) (*sql.Rows, error)
	Update(en db.Entity) error
}

type Repository struct {
	conn db.Connection
}

func NewRepository(connection db.Connection) Repository {
	return Repository{
		conn: connection,
	}
}
