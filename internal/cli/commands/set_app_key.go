package commands

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func SetAppKey() *cobra.Command {
	return &cobra.Command{
		Use:   "sak",
		Short: "Set app key",
		Run: func(cmd *cobra.Command, args []string) {
			randomNumber := rand.Intn(100000)
			hash := sha256.New()
			hash.Write([]byte(strconv.Itoa(randomNumber)))
			hashedBytes := hash.Sum(nil)
			hashString := hex.EncodeToString(hashedBytes)

			err := modifyEnvFile("APP_KEY", hashString)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println("App key set")
		},
	}
}
