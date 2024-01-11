package cli_test

import (
	"testing"

	"github.com/andresmeireles/speaker/internal/cli"
)

func TestCommands(t *testing.T) {
	// act
	cli.Commands()

	// assert
	if t.Failed() {
		t.Fatal("expected nil, got error")
	}
}
