package repository

import (
	"fmt"
	"strings"

	"github.com/andresmeireles/speaker/internal/db"
	"github.com/andresmeireles/speaker/internal/db/entity"
)

func Update[T entity.Entity](en T) error {
	db, err := db.GetDB()

	if err != nil {
		return err
	}

	defer db.Close()

	keys, _, values := Split(en)
	sets := ""

	for _, val := range strings.Split(keys, ",") {
		sets = sets + fmt.Sprintf("%s = ?", val)
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = ?", en.Table(), sets)

	values = append(values, en.GetId())

	_, err = db.Exec(query, values...)

	if err != nil {
		return err
	}

	return nil
}
