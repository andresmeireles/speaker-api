package auxcmd

import (
	"database/sql"
	"os"

	"github.com/andresmeireles/speaker/internal/db"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
)

type Migration interface {
	Up() error
	Down() error
}

type migration struct {
	migrate *migrate.Migrate
}

func NewMigration(conn db.Connection) Migration {
	migrationSource := "file://" + os.Getenv("DB_MIGRATIONS_PATH")
	drive := os.Getenv("DB_DRIVER")
	db, err := conn.GetDB()

	if err != nil {
		panic(err)
	}

	driver, err := getDrive(drive, db)
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		migrationSource,
		drive,
		driver,
	)
	if err != nil {
		panic(err)
	}

	return migration{m}
}

func getDrive(drive string, conn *sql.DB) (database.Driver, error) {
	err := conn.Ping()
	if err != nil {
		panic("error on connection " + err.Error())
	}

	switch drive {
	case "postgres":
		return postgres.WithInstance(conn, &postgres.Config{})
	case "sqlite3":
		return sqlite3.WithInstance(conn, &sqlite3.Config{})
	case "sqlite":
		return sqlite3.WithInstance(conn, &sqlite3.Config{})
	default:
		panic("driver " + drive + " not supported")
	}
}

func (m migration) Up() error {
	return m.Up()
}

func (m migration) Down() error {
	return m.Down()
}
