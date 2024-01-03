package main

import (
	"os"

	"github.com/andresmeireles/speaker/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	envErr := godotenv.Load(dir + "/.env")
	if envErr != nil {
		panic(envErr)
	}

	router.Run(os.Getenv("APP_PORT"))
}
