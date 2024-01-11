package user

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/db/repository"
	"github.com/andresmeireles/speaker/internal/tools/servicelocator"
)

type UserRepository struct {
	repository repository.Repository[entity.User]
}

func (u UserRepository) New(s servicelocator.ServiceLocator) any {
	return UserRepository{
		repository: servicelocator.Get[repository.Repository[entity.User]](s),
	}
}

func (u UserRepository) Add(user entity.User) error {
	return u.repository.Add(user)
}

func (u UserRepository) GetByEmail(email string) (entity.User, error) {
	row, err := u.repository.SingleQuery(
		"SELECT * FROM users WHERE email = $1 LIMIT 1",
		email,
	)
	if err != nil {
		slog.Error("error querying user", err)

		return entity.User{}, err
	}

	user := new(entity.User)
	if err := row.Scan(&user.Id, &user.Name, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			slog.Info("user not found", "email", email)

			return entity.User{}, fmt.Errorf("user with email %s not found", email)
		}

		slog.Error("error scanning ", err)

		return entity.User{}, err
	}

	return *user, nil
}

func (r UserRepository) GetById(id int) (entity.User, error) {
	row, err := r.repository.GetById(id)
	if err != nil {
		return entity.User{}, err
	}

	user := new(entity.User)
	if err = row.Scan(&user.Id, &user.Name, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, fmt.Errorf("user with id %d not found", id)
		}

		return entity.User{}, err
	}

	return *user, nil
}
