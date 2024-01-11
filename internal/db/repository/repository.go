package repository

import (
	"github.com/andresmeireles/speaker/internal/db"
	"github.com/andresmeireles/speaker/internal/db/entity"
	servicelocator "github.com/andresmeireles/speaker/internal/tools/servicelocator"
)

type Repository[T entity.Entity] struct {
	conn db.Connection
}

func (r Repository[T]) New(sl servicelocator.ServiceLocator) any {
	conn := servicelocator.Get[db.Connection](sl)

	return Repository[T]{conn}
}
