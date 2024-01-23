package user

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/andresmeireles/speaker/internal/repository"
)

type UserRepository interface {
	Add(user User) error
	GetById(id int) (User, error)
	GetByEmail(email string) (User, error)
}

type Repository struct {
	repository repository.Repository
}

func NewRepository(repository repository.Repository) Repository {
	return Repository{
		repository: repository,
	}
}

func (u Repository) Add(user User) error {
	return u.repository.Add(user)
}

func (u Repository) GetByEmail(email string) (User, error) {
	row, err := u.repository.SingleQuery(
		"SELECT * FROM users WHERE email = $1 LIMIT 1",
		email,
	)
	if err != nil {
		slog.Error("error querying user", err)

		return User{}, err
	}

	user := new(User)
	if err := row.Scan(&user.Id, &user.Name, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			slog.Info("user not found", "email", email)

			return User{}, fmt.Errorf("user with email %s not found", email)
		}

		slog.Error("error scanning ", err)

		return User{}, err
	}

	return *user, nil
}

func (r Repository) GetById(id int) (User, error) {
	user := new(User)
	row, err := r.repository.GetById(user.Table(), id)

	if err != nil {
		return User{}, err
	}

	if err = row.Scan(&user.Id, &user.Name, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return User{}, fmt.Errorf("user with id %d not found", id)
		}

		return User{}, err
	}

	return *user, nil
}
