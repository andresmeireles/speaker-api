package logger

import (
	"log/slog"
	"os"

	"github.com/andresmeireles/speaker/internal/cli/commands"
)

func createHandler() *slog.TextHandler {
	handlerOptions := slog.HandlerOptions{
		AddSource: true,
	}
	isDev := os.Getenv("APP_MODE") == "dev"
	if isDev {
		root, _ := commands.Root()
		file, err := os.OpenFile(root+"/var/logger.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			panic(err)
		}

		return slog.NewTextHandler(file, &handlerOptions)
	}

	return slog.NewTextHandler(os.Stdout, &handlerOptions)
}

func Logger() {
	handler := createHandler()
	logger := slog.New(handler)
	slog.SetDefault(logger)
}

func Error(values ...any) {
	Logger()
	slog.Error("Some Error Message", values)
}

func Info(values ...any) {
	Logger()
	slog.Info("Some Info Message", values)
}
