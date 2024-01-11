package auth

import (
	"database/sql"
	"fmt"

	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/db/repository"
	"github.com/andresmeireles/speaker/internal/tools/servicelocator"
)

type AuthRepository struct {
	repository repository.Repository[entity.Auth]
}

func (r AuthRepository) New(s servicelocator.ServiceLocator) any {
	re := servicelocator.Get[repository.Repository[entity.Auth]](s)

	return AuthRepository{
		repository: re,
	}
}

func (r AuthRepository) GetByHash(hash string) (entity.Auth, error) {
	auth := new(entity.Auth)
	query := "SELECT * FROM auths WHERE hash = $1 LIMIT 1"
	row, err := r.repository.SingleQuery(query, hash)

	if err != nil {
		return entity.Auth{}, err
	}

	if err := row.Scan(&auth.Id, &auth.UserId, &auth.Hash, &auth.Expired); err != nil {
		if err == sql.ErrNoRows {
			return entity.Auth{}, fmt.Errorf("auth with hash %s not found", hash)
		}

		return entity.Auth{}, err
	}

	return *auth, nil
}

func (a AuthRepository) AuthCodeByUser(authCode string, userId int) (*entity.AuthCode, error) {
	var code entity.AuthCode

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

func (r AuthRepository) Add(auth entity.Auth) error {
	return r.repository.Add(auth)
}

func (r AuthRepository) GetById(id int) (*entity.Auth, error) {
	auth := new(entity.Auth)
	authRow, err := r.repository.GetById(id)

	if err != nil {
		return nil, err
	}

	if err := authRow.Scan(&auth.Id, &auth.User, &auth.Hash, &auth.Expired); err != nil {
		return nil, err
	}

	return auth, nil
}

func (r AuthRepository) GetAll() ([]entity.Auth, error) {
	auth := new(entity.Auth)
	auths := make([]entity.Auth, 0)
	authRows, err := r.repository.GetAll()

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

func (a AuthRepository) Update(auth entity.Auth) error {
	return a.repository.Update(auth)
}

func (a AuthRepository) Delete(auth entity.Auth) error {
	return a.repository.Delete(auth)
}
