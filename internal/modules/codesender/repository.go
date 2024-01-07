package codesender

import (
	"database/sql"
	"fmt"

	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/db/repository"
	"github.com/andresmeireles/speaker/internal/modules/user"
)

type AuthCodeRepository struct {
	userRepository user.UserRepository
}

func (a AuthCodeRepository) Add(authCode entity.AuthCode) error {
	return repository.Add[entity.AuthCode](authCode)
}

func (a AuthCodeRepository) GetById(authCodeId int) (entity.AuthCode, error) {
	row, err := repository.GetById[entity.AuthCode](authCodeId)
	if err != nil {
		return entity.AuthCode{}, err
	}

	var authCode entity.AuthCode
	if err := row.Scan(&authCode.Id, &authCode.Code, &authCode.User, &authCode.ExpiresAt); err != nil {
		return entity.AuthCode{}, err
	}

	return authCode, nil
}

func (a AuthCodeRepository) GetAll() ([]entity.AuthCode, error) {
	codes := make([]entity.AuthCode, 0)
	rows, err := repository.GetAll[entity.AuthCode]()

	if err != nil {
		return codes, err
	}

	for rows.Next() {
		var authCode entity.AuthCode

		if err := rows.Scan(&authCode.Id, &authCode.Code, &authCode.User, &authCode.ExpiresAt); err != nil {
			return nil, err
		}

		codes = append(codes, authCode)
	}

	return codes, nil
}

func (a AuthCodeRepository) GetByCode(code string) (entity.AuthCode, error) {
	query := "SELECT * FROM auth_codes WHERE code = $1 LIMIT 1"
	row, err := repository.SingleQuery(query, code)

	if err != nil {
		return entity.AuthCode{}, err
	}

	authCode := new(entity.AuthCode)
	if err = row.Scan(&authCode.Id, &authCode.Code, &authCode.ExpiresAt, &authCode.UserId); err != nil {
		if err == sql.ErrNoRows {
			return entity.AuthCode{}, fmt.Errorf("auth code with code %s not found", code)
		}

		return entity.AuthCode{}, err
	}

	user, err := a.userRepository.GetById(authCode.UserId)
	if err != nil {
		return entity.AuthCode{}, err
	}

	authCode.User = user

	return *authCode, nil
}

func (a AuthCodeRepository) Update(authCode entity.AuthCode) error {
	return repository.Update(authCode)
}

func (a AuthCodeRepository) Delete(authCode entity.AuthCode) error {
	return repository.Delete(authCode)
}
