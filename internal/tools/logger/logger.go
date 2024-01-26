package logger

import (
	"log/slog"
	"os"

	"github.com/andresmeireles/speaker/internal/cli/commands"
)

const PERMISSIONS = 0644

func createHandler() *slog.JSONHandler {
	handlerOptions := slog.HandlerOptions{}
	isDev := os.Getenv("APP_MODE") == "dev"

	if isDev {
		root, _ := commands.Root()
		file, err := os.OpenFile(root+"/var/logger.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, PERMISSIONS)

		if err != nil {
			panic(err)
		}

		// return slog.NewTextHandler(file, &handlerOptions)
		return slog.NewJSONHandler(file, &handlerOptions)
	}

	return slog.NewJSONHandler(os.Stdout, &handlerOptions)
	// return slog.NewTextHandler(os.Stdout, &handlerOptions)
}

func Logger() {
	handler := createHandler()
	logger := slog.New(handler)
	slog.SetDefault(logger)
}
