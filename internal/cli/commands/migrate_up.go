package commands

import (
	"github.com/andresmeireles/speaker/internal/database"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/spf13/cobra"
)

func MigrateUp() *cobra.Command {
	return &cobra.Command{
		Use:   "mup",
		Short: "Migrate up",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := database.GetDB()

			if err != nil {
				panic(err)
			}

			driver, err := postgres.WithInstance(conn, &postgres.Config{})

			if err != nil {
				panic(err)
			}
			migration, err := migrate.NewWithDatabaseInstance(
				"file://../internal/database/migrations",
				"postgres",
				driver,
			)

			if err != nil {
				panic(err)
			}

			migration.Up()
		},
	}
}
