package cli_test

import (
	"testing"

	"github.com/andresmeireles/speaker/internal/cli"
	"github.com/andresmeireles/speaker/internal/tools/servicelocator"
)

func TestCommands(t *testing.T) {
	// arrange
	sl := servicelocator.NewServiceLocator()

	// act
	cli.Commands(*sl)

	// assert
	if t.Failed() {
		t.Fatal("expected nil, got error")
	}
}
