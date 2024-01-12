package repository

import (
	"fmt"
	"strings"

	"github.com/andresmeireles/speaker/internal/db"
)

func (r Repository) Update(en db.Entity) error {
	db, err := r.conn.GetDB()
	if err != nil {
		return err
	}
	defer db.Close()

	keys, _, values := r.split(en)
	sets := ""
	lastParam := len(values) + 1

	for key, val := range strings.Split(keys, ",") {
		sets = sets + fmt.Sprintf("%s = $%d,", val, key+1)
	}

	sets = strings.Trim(sets, ",")
	query := fmt.Sprintf(
		"UPDATE %s SET %s WHERE id = $%d",
		en.Table(),
		sets,
		lastParam,
	)

	values = append(values, en.GetId())

	_, err = db.Exec(query, values...)
	if err != nil {
		return err
	}

	return nil
}
