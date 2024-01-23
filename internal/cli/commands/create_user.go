package commands

import (
	"os"

	"github.com/andresmeireles/speaker/internal/user"
	"github.com/spf13/cobra"
)

func CreateUser(userRepository user.UserRepository) *cobra.Command {
	var name, email string

	command := &cobra.Command{
		Use:   "cuser",
		Short: "Create user",
		Run: func(cmd *cobra.Command, args []string) {
			if name == "" {
				os.Stderr.Write([]byte("Name is required\n"))
				os.Exit(1)
			}
			if email == "" {
				os.Stderr.Write([]byte("Email is required\n"))
				os.Exit(1)
			}

			user := user.User{
				Name:  name,
				Email: email,
			}
			err := userRepository.Add(user)

			if err != nil {
				os.Stderr.WriteString("Error when creating user\n")
				os.Exit(1)
			}

			os.Stdout.WriteString("User created\n")
		},
	}

	command.Flags().StringVarP(&name, "name", "n", "", "Name")
	command.Flags().StringVarP(&email, "email", "e", "", "Email")

	return command
}
