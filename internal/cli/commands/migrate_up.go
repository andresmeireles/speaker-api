package commands

import (
	"os"

	"github.com/spf13/cobra"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func migrationSource() string {
	return "file://" + os.Getenv("DB_MIGRATIONS_PATH")
}

// func migrationSetup() *migrate.Migrate {
// 	conn, err := db.GetDB()
// 	if err != nil {
// 		panic(err)
// 	}

// 	drive := os.Getenv("DB_DRIVER")
// 	driver, err := getDrive(drive, conn)

// 	if err != nil {
// 		panic(err)
// 	}

// 	migrationSource := migrationSource()
// 	migration, err := migrate.NewWithDatabaseInstance(
// 		migrationSource,
// 		drive,
// 		driver,
// 	)

// 	if err != nil {
// 		fmt.Println("error!")
// 		panic(err)
// 	}

// 	return migration
// }

// func getDrive(drive string, conn *sql.DB) (migrateDatabase.Driver, error) {
// 	switch drive {
// 	case "postgres":
// 		return postgres.WithInstance(conn, &postgres.Config{})
// 	case "sqlite3":
// 		return sqlite3.WithInstance(conn, &sqlite3.Config{})
// 	default:
// 		panic("driver " + drive + " not supported")
// 	}
// }

func MigrateUp() *cobra.Command {
	return &cobra.Command{
		Use:   "mup",
		Short: "Migrate up",
		Run: func(cmd *cobra.Command, args []string) {
			// migration := migrationSetup()

			// err := migration.Up()

			// if err == migrate.ErrNoChange {
			// 	fmt.Println("Nothing to migrate")
			// 	return
			// }

			// if err != nil {
			// 	fmt.Println(err)
			// 	os.Exit(1)
			// }

			// fmt.Println("Migration up done")
		},
	}
}
