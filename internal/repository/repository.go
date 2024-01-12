package repository

import (
	"github.com/andresmeireles/speaker/internal/db"
	servicelocator "github.com/andresmeireles/speaker/internal/tools/servicelocator"
)

type Repository[T db.Entity] struct {
	conn db.Connection
}

func (r Repository[T]) New(sl servicelocator.ServiceLocator) any {
	conn := servicelocator.Get[db.Connection](sl)

	return Repository[T]{conn}
}
