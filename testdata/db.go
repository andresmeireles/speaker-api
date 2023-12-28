package testdata

import (
	"os"
	"testing"

	"github.com/andresmeireles/speaker/internal/cli/commands"
)

func SetupDatabase(m *testing.M) {
	SetCredentials()

	TeardownLocalDB()
	err := SetupLocalDB()

	if err != nil {
		TeardownLocalDB()
		panic(err)
	}

	m.Run()

	TeardownLocalDB()
}

func SetupLocalDB() error {
	return commands.MigrateUp().Execute()
}

func SetCredentials() {
	os.Setenv("DB_DRIVER", "postgres")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5433")
	os.Setenv("DB_USERNAME", "speaker")
	os.Setenv("DB_PASSWORD", "speaker")
}

func TeardownLocalDB() {
	commands.MigrateDown().Execute()
}
