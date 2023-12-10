package testdata

import (
	"os"
	"testing"

	"github.com/andresmeireles/speaker/internal/cli/commands"
)

func SetupDatabase(m *testing.M) {
	SetupLocalDB()
	m.Run()
	TeardownLocalDB()
}

func SetupLocalDB() {
	os.Setenv("DB_DRIVER", "sqlite3")

	commands.MigrateUp().Execute()
}

func TeardownLocalDB() {
	os.Setenv("DB_DRIVER", "sqlite3")

	commands.MigrateDown().Execute()
}
