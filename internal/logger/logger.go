package logger

import (
	"log/slog"
	"os"

	"github.com/andresmeireles/speaker/internal/cli/commands"
)

func logger() {
	root, _ := commands.Root()
	file, err := os.OpenFile(root+"/var/logger.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	handler := slog.NewTextHandler(file, nil)
	logger := slog.New(handler)
	slog.SetDefault(logger)
}

func Error(values ...any) {
	logger()
	slog.Error("Some Error Message", values)
}

func Info(values ...any) {
	logger()
	slog.Info("Some Info Message", values)
}
