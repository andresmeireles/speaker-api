package cli

import (
	"fmt"
	"os"

	"github.com/andresmeireles/speaker/internal/cli/commands"
	"github.com/spf13/cobra"
)

func Commands() {
	cmd := &cobra.Command{}

	cmd.AddCommand(
		commands.MigrateUp(),
		commands.MigrateDown(),
		commands.SetAppKey(),
	)

	err := cmd.Execute()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
