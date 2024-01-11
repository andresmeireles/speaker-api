package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrateDown() *cobra.Command {
	return &cobra.Command{
		Use:   "mud",
		Short: "Migrate down",
		Long:  "Undo all migrations",
		Run: func(cmd *cobra.Command, args []string) {
			migration := migrationSetup()

			err := migration.Down()

			if err == migrate.ErrNoChange {
				fmt.Println("Nothing to rollback")
				return
			}

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println("Migration down done")
		},
	}
}
