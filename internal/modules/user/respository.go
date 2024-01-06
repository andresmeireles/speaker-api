package user

import (
	"log/slog"

	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/db/repository"
)

type UserRepository struct{}

func (u UserRepository) UserByEmail(email string) (entity.User, error) {
	row := repository.SingleQuery("SELECT * FROM users WHERE email = $1 LIMIT 1", email)
	if row.Err() != nil {
		return entity.User{}, row.Err()
	}

	user := new(entity.User)
	if err := row.Scan(&user.Id, &user.Email, &user.Name); err != nil {
		slog.Error("error scanning user", err)

		return entity.User{}, err
	}

	return *user, nil
}
