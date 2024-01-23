package main

import (
	"os"

	"github.com/andresmeireles/speaker/internal"
	"github.com/andresmeireles/speaker/internal/tools/logger"
	"github.com/andresmeireles/speaker/internal/tools/servicelocator"
	"github.com/andresmeireles/speaker/internal/web/router"
	"github.com/joho/godotenv"
)

func main() {
	os.Setenv("TZ", "America/Sao_Paulo")
	dir, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	mode := os.Getenv("APP_MODE")
	if mode == "" {
		mode = "dev"
	}

	if mode == "dev" {
		if envErr := godotenv.Load(dir + "/.env"); envErr != nil {
			panic(envErr)
		}
	}

	logger.Logger()

	sl := servicelocator.NewServiceLocator()
	internal.DIContainer(sl)
	router.Run(os.Getenv("APP_PORT"), sl)
}
