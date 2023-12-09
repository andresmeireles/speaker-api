package repository

import "github.com/andresmeireles/speaker/internal/db/entity"

type Repository[T entity.Entity] interface {
	Add(T) error
	GetById(int) (*T, error)
	GetAll() []T
	Update(T) error
	Delete(T) error
}
