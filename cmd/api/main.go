package main

import (
	"os"

	"github.com/andresmeireles/speaker/internal/logger"
	"github.com/andresmeireles/speaker/internal/web/router"
	"github.com/joho/godotenv"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	mode := os.Getenv("APP_MODE")
	if mode == "" {
		mode = "dev"
	}

	if mode == "dev" {
		envErr := godotenv.Load(dir + "/.env")
		if envErr != nil {
			panic(envErr)
		}
	}

	logger.Logger()

	router.Run(os.Getenv("APP_PORT"))
}
