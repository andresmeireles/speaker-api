package codesender

import (
	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/db/repository"
)

type AuthCodeRepository struct{}

func (a AuthCodeRepository) Add(authCode entity.AuthCode) error {
	return repository.Add[entity.AuthCode](authCode)
}

func (a AuthCodeRepository) GetById(authCodeId int) (entity.AuthCode, error) {
	row := repository.GetById[entity.AuthCode](authCodeId)

	if row.Err() != nil {
		return entity.AuthCode{}, row.Err()
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

func (a AuthCodeRepository) Update(authCode entity.AuthCode) error {
	return repository.Update(authCode)
}

func (a AuthCodeRepository) Delete(authCode entity.AuthCode) error {
	return repository.Delete(authCode)
}
