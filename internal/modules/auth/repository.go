package auth

import (
	"database/sql"
	"fmt"

	"github.com/andresmeireles/speaker/internal/db"
	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/db/repository"
)

type AuthRepository struct{}

func (a AuthRepository) GetByHash(hash string) (entity.Auth, error) {
	db, err := db.GetDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var auth entity.Auth
	query := "SELECT * FROM auths WHERE hash = ?"
	row := db.QueryRow(query, hash)

	if err := row.Scan(&auth); err != nil {
		if err == sql.ErrNoRows {
			return entity.Auth{}, fmt.Errorf("auth with hash %s not found", hash)
		}
	}

	return auth, nil
}

func (a AuthRepository) AuthCodeByUser(authCode string, userId int) (*entity.AuthCode, error) {
	db, err := db.GetDB()
	if err != nil {
		panic(err) // RESOLVER ISSO NO FUTURO
	}
	defer db.Close()

	var code entity.AuthCode
	query := "SELECT * FROM auth_code WHERE code = ? AND user_id = ?"
	row := db.QueryRow(query, authCode, userId)
	if err := row.Scan(&code); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
	}

	return &code, nil
}

func (a AuthRepository) Add(auth entity.Auth) error {
	return repository.Add(auth)
}

func (a AuthRepository) GetById(id int) (*entity.Auth, error) {
	auth := new(entity.Auth)
	authRow := repository.GetById[entity.Auth](id)

	if err := authRow.Scan(&auth.Id, &auth.User, &auth.Hash, &auth.Expired); err != nil {
		return nil, err
	}

	return auth, nil
}

func (a AuthRepository) GetAll() ([]entity.Auth, error) {
	auth := new(entity.Auth)
	auths := make([]entity.Auth, 0)
	authRows, err := repository.GetAll[entity.Auth]()

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

func (a AuthRepository) Update(auth entity.Auth) error {
	return repository.Update(auth)
}

func (a AuthRepository) Delete(auth entity.Auth) error {
	return repository.Delete(auth)
}
