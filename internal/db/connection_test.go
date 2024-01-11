package db_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/andresmeireles/speaker/internal/db"
	"github.com/andresmeireles/speaker/testdata"
)

func TestGetDB(t *testing.T) {
	os.Setenv("DB_DRIVER", "sqlite3")
	os.Setenv("SMTP_USER", "andre.meireles@gmail.com")

	db, err := testdata.GetService[db.Connection]().GetDB()
	if err != nil {
		t.Fatalf("expected nil, got %s", err)
	}

	if reflect.TypeOf(db).String() != "*sql.DB" {
		t.Fatalf("expected *db.DB, got %s", reflect.TypeOf(db).String())
	}
}

func TestGetDBPostgres(t *testing.T) {
	os.Setenv("DB_DRIVER", "postgres")
	os.Setenv("SMTP_USER", "andre.meireles@gmail.com")

	db, err := testdata.GetService[db.Connection]().GetDB()
	if err != nil {
		t.Fatalf("expected nil, got %s", err)
	}

	if reflect.TypeOf(db).String() != "*sql.DB" {
		t.Fatalf("expected *db.DB, got %s", reflect.TypeOf(db).String())
	}
}
