package main

import (
	"log/slog"
	"os"
	"ref_system/migrations"
)

func main() {
	loggerHandler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(loggerHandler)
	slog.SetDefault(logger)

	if err := migrations.Up(logger); err != nil {
		logger.Error("failed to apply migrations", slog.String("error", err.Error()))
	}
}
