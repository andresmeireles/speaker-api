package user

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/andresmeireles/speaker/internal/repository"
	"github.com/andresmeireles/speaker/internal/tools/servicelocator"
)

type UserRepository struct {
	repository repository.Repository[User]
}

func (u UserRepository) New(s servicelocator.ServiceLocator) any {
	return UserRepository{
		repository: servicelocator.Get[repository.Repository[User]](s),
	}
}

func (u UserRepository) Add(user User) error {
	return u.repository.Add(user)
}

func (u UserRepository) GetByEmail(email string) (User, error) {
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

func (r UserRepository) GetById(id int) (User, error) {
	row, err := r.repository.GetById(id)
	if err != nil {
		return User{}, err
	}

	user := new(User)
	if err = row.Scan(&user.Id, &user.Name, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return User{}, fmt.Errorf("user with id %d not found", id)
		}

		return User{}, err
	}

	return *user, nil
}
