package main

import (
	"os"

	"github.com/andresmeireles/speaker/internal"
	"github.com/andresmeireles/speaker/internal/cli"
	"github.com/andresmeireles/speaker/internal/tools/servicelocator"
	"github.com/joho/godotenv"
)

func main() {
	mode := os.Getenv("APP_MODE")
	if mode == "" {
		mode = "dev"
	}

	isDev := mode == "dev"
	dir, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	if isDev {
		if envErr := godotenv.Load(dir + "/.env"); envErr != nil {
			panic(envErr)
		}
	}

	sl := servicelocator.NewServiceLocator()
	internal.DIContainer(sl)

	cli.Commands(sl)
}
