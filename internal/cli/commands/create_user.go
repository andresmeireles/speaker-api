package commands

import (
	"fmt"
	"os"

	"github.com/andresmeireles/speaker/internal/user"
	"github.com/spf13/cobra"
)

func CreateUser() *cobra.Command {
	var name, email string

	command := &cobra.Command{
		Use:   "cuser",
		Short: "Create user",
		Run: func(cmd *cobra.Command, args []string) {
			if name == "" {
				fmt.Println("Name is required")
				os.Exit(1)
			}
			if email == "" {
				fmt.Println("Email is required")
				os.Exit(1)
			}

			userRepository := user.UserRepository{}
			user := user.User{
				Name:  name,
				Email: email,
			}
			err := userRepository.Add(user)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println("User created")
		},
	}

	command.Flags().StringVarP(&name, "name", "n", "", "Name")
	command.Flags().StringVarP(&email, "email", "e", "", "Email")

	return command
}
