package commands

import (
	"fmt"
	"os"

	"github.com/andresmeireles/speaker/internal/db"
	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/spf13/cobra"
)

func ListUser() *cobra.Command {
	return &cobra.Command{
		Use:   "luser",
		Short: "List users",
		Run: func(cmd *cobra.Command, args []string) {
			db, err := db.GetDB()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			r, e := db.Query("SELECT name FROM users")

			if e != nil {
				fmt.Fprintln(os.Stderr, e)
			}

			var u []entity.User

			for r.Next() {
				r.Scan(u)
			}

			for _, v := range u {
				fmt.Println(v.Name)
			}
		},
	}
}
