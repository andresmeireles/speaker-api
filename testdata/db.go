package testdata

import (
	"os"
	"testing"

	"github.com/andresmeireles/speaker/internal/cli/commands"
)

func Setup(m *testing.M) {
	setupLocalDB()
	m.Run()
	teardownLocalDB()
}

func setupLocalDB() {
	os.Setenv("DB_DRIVER", "sqlite3")

	commands.MigrateUp().Execute()
}

func teardownLocalDB() {
	os.Setenv("DB_DRIVER", "sqlite3")

	commands.MigrateDown().Execute()
}
