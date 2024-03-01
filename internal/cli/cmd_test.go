package cli_test

import (
	"testing"

	"github.com/andresmeireles/speaker/internal/cli"
	"github.com/andresmeireles/speaker/testdata/mocks"
)

func TestCommands(t *testing.T) {
	// arrange
	sl := mocks.ServiceLocatorservicelocator{}
	sl.EXPECT().Get("user.UserRepository").Return(&mocks.UserRepositoryuser{})
	sl.EXPECT().Get("auxcmd.Migration").Return(&mocks.Migrationauxcmd{})

	// act
	cli.Commands(&sl)

	// assert
	if t.Failed() {
		t.Fatal("expected nil, got error")
	}
}
