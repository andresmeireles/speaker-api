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
	os.Setenv("APP_KEY", "e7bca8464289691d92f60271")
	os.Setenv("SMTP_HOST", "smtp.server")
	os.Setenv("SMTP_PORT", "465")
	os.Setenv("SMTP_USER", "email@gmail.com")
	os.Setenv("SMTP_PASSWORD", "password")
	os.Setenv("APP_MODE", "test")

	// db envs
	root, _ := commands.Root()
	os.Setenv("DB_MIGRATIONS_PATH", root+"/build/migrations")
}

func TeardownLocalDB() {
	commands.MigrateDown().Execute()
}
