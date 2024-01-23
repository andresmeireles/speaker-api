package auth

import (
	"database/sql"
	"fmt"

	"github.com/andresmeireles/speaker/internal/codesender"
	"github.com/andresmeireles/speaker/internal/repository"
)

type AuthRepository struct {
	repository repository.Repository
}

func NewRepository(repository repository.Repository) AuthRepository {
	return AuthRepository{
		repository: repository,
	}
}

func (r AuthRepository) GetByHash(hash string) (Auth, error) {
	auth := new(Auth)
	query := "SELECT * FROM auths WHERE hash = $1 LIMIT 1"
	row, err := r.repository.SingleQuery(query, hash)

	if err != nil {
		return Auth{}, err
	}

	if err := row.Scan(&auth.Id, &auth.UserId, &auth.Hash, &auth.Expired); err != nil {
		if err == sql.ErrNoRows {
			return Auth{}, fmt.Errorf("auth with hash %s not found", hash)
		}

		return Auth{}, err
	}

	return *auth, nil
}

func (a AuthRepository) AuthCodeByUser(authCode string, userId int) (*codesender.AuthCode, error) {
	var code codesender.AuthCode

	query := "SELECT * FROM auth_code WHERE code = ? AND user_id = ?"
	row, err := a.repository.SingleQuery(query, authCode, userId)

	if err != nil {
		return nil, err
	}

	if err := row.Scan(&code); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
	}

	return &code, nil
}

func (r AuthRepository) Add(auth Auth) error {
	return r.repository.Add(auth)
}

func (r AuthRepository) GetById(id int) (*Auth, error) {
	auth := new(Auth)
	authRow, err := r.repository.GetById(auth.Table(), id)

	if err != nil {
		return nil, err
	}

	if err := authRow.Scan(&auth.Id, &auth.User, &auth.Hash, &auth.Expired); err != nil {
		return nil, err
	}

	return auth, nil
}

func (r AuthRepository) GetAll() ([]Auth, error) {
	auth := new(Auth)
	auths := make([]Auth, 0)
	authRows, err := r.repository.GetAll(auth.Table())

	if err != nil {
		return nil, err
	}

	for authRows.Next() {
		if err := authRows.Scan(&auth.Id, &auth.User, &auth.Hash, &auth.Expired); err != nil {
			return nil, err
		}

		auths = append(auths, *auth)
	}

	return auths, nil
}

func (r AuthRepository) ExpireTokenByUserId(userId int) error {
	query := "UPDATE auths SET expired = true WHERE user_id = $1 AND expired = false"
	_, err := r.repository.Query(query, userId)

	return err
}

func (a AuthRepository) Update(auth Auth) error {
	return a.repository.Update(auth)
}

func (a AuthRepository) Delete(auth Auth) error {
	return a.repository.Delete(auth)
}
