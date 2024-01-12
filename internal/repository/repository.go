package repository

import (
	"github.com/andresmeireles/speaker/internal/db"
	servicelocator "github.com/andresmeireles/speaker/internal/tools/servicelocator"
)

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
