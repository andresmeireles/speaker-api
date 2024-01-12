package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/andresmeireles/speaker/internal/tools/servicelocator"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type Connection struct{}

func NewConnection() Connection {
	return Connection{}
}

func (c Connection) New(_ servicelocator.ServiceLocator) any {
	return Connection{}
}

func (c Connection) GetDB() (*sql.DB, error) {
	driver := os.Getenv("DB_DRIVER")
	connectionString, err := queryStringByDrive(driver)

	if err != nil {
		return nil, err
	}

	return sql.Open(driver, connectionString)
}

func queryStringByDrive(driver string) (string, error) {
	switch driver {
	case "postgres":
		return postgres(), nil
	case "sqlite3":
		return sqlite(), nil
	case "sqlite":
		return sqlite(), nil
	default:
		return "", errors.New("driver " + driver + " not supported")
	}
}

func postgres() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
}

func sqlite() string {
	// create a sqlite3 in memory db connection string
	// return ":memory:"
	return "sdb.db"
}
