package commands

import (
	"fmt"
	"os"

	"github.com/andresmeireles/speaker/internal/db"
	"github.com/spf13/cobra"
)

func CreateUser() *cobra.Command {
	return &cobra.Command{
		Use:   "cuser",
		Short: "Create user",
		Run: func(cmd *cobra.Command, args []string) {
			db, err := db.GetDB()

			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}

			db.Query("INSERT INTO users (name) VALUES ('andres')")

			defer db.Close()

			fmt.Println("User created")
		},
	}
}
