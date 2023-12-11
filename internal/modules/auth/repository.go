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

func (a AuthRepository) Add(auth entity.Auth) error {
	return repository.Add(auth)
}

func (a AuthRepository) GetById(id int) (*entity.Auth, error) {
	return repository.GetById[entity.Auth](id, entity.Auth{})
}

func (a AuthRepository) GetAll() []entity.Auth {
	return repository.GetAll[entity.Auth](entity.Auth{})
}

func (a AuthRepository) Update(auth entity.Auth) error {
	return repository.Update(auth)
}

func (a AuthRepository) Delete(auth entity.Auth) error {
	return repository.Delete(auth)
}
