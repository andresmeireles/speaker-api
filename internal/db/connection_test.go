package db_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/andresmeireles/speaker/internal/db"
)

func TestDB(t *testing.T) {
	os.Setenv("DB_DRIVER", "sqlite3")

	conn := db.Connection{}

	t.Run("Get DB", func(t *testing.T) {
		os.Setenv("DB_DRIVER", "sqlite3")

		db, err := conn.GetDB()

		if err != nil {
			t.Fatalf("expected nil, got %s", err)
		}

		if reflect.TypeOf(db).String() != "*sql.DB" {
			t.Fatalf("expected *db.DB, got %s", reflect.TypeOf(db).String())
		}
	})
}
