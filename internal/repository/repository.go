package repository

import (
	"database/sql"

	"github.com/andresmeireles/speaker/internal/db"
	servicelocator "github.com/andresmeireles/speaker/internal/tools/servicelocator"
)

type RepositoryInterface interface {
	Add(en db.Entity) error
	Delete(en db.Entity) error
	Update(en db.Entity) error
	GetById(id int) (db.Entity, error)
	GetAll() ([]db.Entity, error)
	Query(query string, values ...any) (*sql.Rows, error)
	SingleQuery(query string, values ...any) (*sql.Row, error)
}

type Repository struct {
	conn db.Connection
}

func NewRepository(connection db.Connection) Repository {
	return Repository{
		conn: connection,
	}
}

func (r Repository) New(sl servicelocator.ServiceLocator) any {
	conn := servicelocator.Get[db.Connection](sl)

	return Repository{conn}
}
