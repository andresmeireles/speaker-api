package commands_test

import (
	"testing"

	"github.com/andresmeireles/speaker/internal/cli/commands"
)

func TestGetRoot(t *testing.T) {
	// act
	root, err := commands.Root()

	// assert
	if err != nil {
		t.Fatalf("expected nil, got %s", err)
	}

	if root == "" {
		t.Fatalf("expected non empty string, got %s", root)
	}
}
