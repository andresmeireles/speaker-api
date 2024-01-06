package user

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/db/repository"
)

type UserRepository struct{}

func (u UserRepository) UserByEmail(email string) (entity.User, error) {
	row, err := repository.SingleQuery("SELECT * FROM users WHERE email = $1 LIMIT 1", email)
	if err != nil {
		return entity.User{}, err
	}

	user := new(entity.User)
	if err := row.Scan(&user.Id, &user.Email, &user.Name); err != nil {
		if err == sql.ErrNoRows {
			slog.Info("user not found", "email", email)

			return entity.User{}, fmt.Errorf("user with email %s not found", email)
		}

		slog.Error("error scanning ", err)

		return entity.User{}, err
	}

	return *user, nil
}
