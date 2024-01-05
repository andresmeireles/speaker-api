package logger

import (
	"log/slog"
	"os"

	"github.com/andresmeireles/speaker/internal/cli/commands"
)

func Logger() {
	root, _ := commands.Root()
	file, err := os.OpenFile(root+"/var/logger.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	handler := slog.NewTextHandler(file, &slog.HandlerOptions{
		AddSource: true,
	})
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
