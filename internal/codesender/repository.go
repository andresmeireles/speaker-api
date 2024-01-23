package codesender

import (
	"database/sql"
	"fmt"

	"github.com/andresmeireles/speaker/internal/repository"
	"github.com/andresmeireles/speaker/internal/tools/servicelocator"
	"github.com/andresmeireles/speaker/internal/user"
)

type CodeSenderRepository interface {
	Add(authCode AuthCode) error
	GetById(authCodeId int) (AuthCode, error)
	GetAll() ([]AuthCode, error)
	GetByCode(code string) (AuthCode, error)
	Update(authCode AuthCode) error
}

type Repository struct {
	repository     repository.Repository
	userRepository user.Repository
}

func NewRepository(repository repository.Repository, userRepository user.Repository) Repository {
	return Repository{
		repository:     repository,
		userRepository: userRepository,
	}
}

func (r Repository) New(s servicelocator.ServiceLocator) any {
	return Repository{
		repository:     servicelocator.Get[repository.Repository](s),
		userRepository: servicelocator.Get[user.Repository](s),
	}
}

func (r Repository) Add(authCode AuthCode) error {
	return r.repository.Add(authCode)
}

func (a Repository) GetById(authCodeId int) (AuthCode, error) {
	authCode := new(AuthCode)
	row, err := a.repository.GetById(authCode.Table(), authCodeId)

	if err != nil {
		return AuthCode{}, err
	}

	if err := row.Scan(&authCode.Id, &authCode.Code, &authCode.User, &authCode.ExpiresAt); err != nil {
		return AuthCode{}, err
	}

	return *authCode, nil
}

func (a Repository) GetAll() ([]AuthCode, error) {
	codes := make([]AuthCode, 0)
	rows, err := a.repository.GetAll(AuthCode{}.Table())

	if err != nil {
		return codes, err
	}

	for rows.Next() {
		var authCode AuthCode

		if err := rows.Scan(&authCode.Id, &authCode.Code, &authCode.User, &authCode.ExpiresAt); err != nil {
			return nil, err
		}

		codes = append(codes, authCode)
	}

	return codes, nil
}

func (a Repository) GetByCode(code string) (AuthCode, error) {
	query := "SELECT * FROM auth_codes WHERE code = $1 LIMIT 1"
	row, err := a.repository.SingleQuery(query, code)

	if err != nil {
		return AuthCode{}, err
	}

	authCode := new(AuthCode)
	if err = row.Scan(&authCode.Id, &authCode.Code, &authCode.ExpiresAt, &authCode.UserId); err != nil {
		if err == sql.ErrNoRows {
			return AuthCode{}, fmt.Errorf("auth code with code %s not found", code)
		}

		return AuthCode{}, err
	}

	user, err := a.userRepository.GetById(authCode.UserId)
	if err != nil {
		return AuthCode{}, err
	}

	authCode.User = user

	return *authCode, nil
}

func (a Repository) Update(authCode AuthCode) error {
	return a.repository.Update(authCode)
}

func (a Repository) Delete(authCode AuthCode) error {
	return a.repository.Delete(authCode)
}
