package codesender

import (
	"database/sql"
	"fmt"

	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/db/repository"
	"github.com/andresmeireles/speaker/internal/modules/user"
	"github.com/andresmeireles/speaker/internal/tools/servicelocator"
)

type AuthCodeRepository struct {
	repository     repository.Repository[entity.AuthCode]
	userRepository user.UserRepository
}

func (r AuthCodeRepository) New(s servicelocator.ServiceLocator) any {
	return AuthCodeRepository{
		repository:     servicelocator.Get[repository.Repository[entity.AuthCode]](s),
		userRepository: servicelocator.Get[user.UserRepository](s),
	}
}

func (r AuthCodeRepository) Add(authCode entity.AuthCode) error {
	return r.repository.Add(authCode)
}

func (a AuthCodeRepository) GetById(authCodeId int) (entity.AuthCode, error) {
	row, err := a.repository.GetById(authCodeId)
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
	rows, err := a.repository.GetAll()

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
	row, err := a.repository.SingleQuery(query, code)

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
	return a.repository.Update(authCode)
}

func (a AuthCodeRepository) Delete(authCode entity.AuthCode) error {
	return a.repository.Delete(authCode)
}
