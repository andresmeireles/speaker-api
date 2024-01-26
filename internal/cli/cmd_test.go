package cli_test

import (
	"testing"

	"github.com/andresmeireles/speaker/internal/cli"
	"github.com/andresmeireles/speaker/testdata/mocks"
)

func TestCommands(t *testing.T) {
	// arrange
	sl := mocks.ServiceLocator{}
	sl.EXPECT().Get("user.UserRepository").Return(&mocks.UserRepository{})
	sl.EXPECT().Get("auxcmd.Migration").Return(&mocks.Migration{})

	// act
	cli.Commands(&sl)

	// assert
	if t.Failed() {
		t.Fatal("expected nil, got error")
	}
}
