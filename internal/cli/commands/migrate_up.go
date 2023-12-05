package commands

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/andresmeireles/speaker/internal/database"
	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/cobra"

	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	migrateDatabase "github.com/golang-migrate/migrate/v4/database"
)

func root() (string, error) {
	dir, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	for {
		goModPath := filepath.Join(dir, "go.mod")
		_, err := os.Stat(goModPath)
		if err == nil {
			return dir, nil
		}
		if !os.IsNotExist(err) {
			return "", err
		}

		// go upwards
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break
		}
		dir = parentDir
	}

	return "", fmt.Errorf("go.mod not found")
}

func migrationSetup() *migrate.Migrate {
	conn, err := database.GetDB()

	if err != nil {
		panic(err)
	}

	drive := os.Getenv("DB_DRIVER")

	driver, err := getDrive(drive, conn)

	if err != nil {
		panic(err)
	}

	root, err := root()

	if err != nil {
		panic(err)
	}

	databasePath := filepath.Join(root, "internal", "database", "migration")

	migration, err := migrate.NewWithDatabaseInstance(
		"file://"+databasePath,
		drive,
		driver,
	)

	if err != nil {
		panic(err)
	}

	return migration
}

func getDrive(drive string, conn *sql.DB) (migrateDatabase.Driver, error) {
	switch drive {
	case "postgres":
		return postgres.WithInstance(conn, &postgres.Config{})
	case "sqlite3":
		return sqlite3.WithInstance(conn, &sqlite3.Config{})
	default:
		panic("driver " + drive + " not supported")
	}
}

func MigrateUp() *cobra.Command {
	return &cobra.Command{
		Use:   "mup",
		Short: "Migrate up",
		Run: func(cmd *cobra.Command, args []string) {
			migration := migrationSetup()

			err := migration.Up()

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println("Migration up done")
			os.Exit(0)
		},
	}
}
