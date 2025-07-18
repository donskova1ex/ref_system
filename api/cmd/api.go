package main

import (
	"log/slog"
	"os"
	"ref_system/internal/config"
	"ref_system/internal/repository"
	"ref_system/internal/router"
	"ref_system/migrations"
	"ref_system/pkg/db"
)

func main() {
	logger := loggerInit()

	logger.Info("Configuration initialization has started")
	cfg, err := config.New()
	if err != nil {
		logger.Error("failed to load config", slog.String("error", err.Error()))
	}
	logger.Info("Configuration initialization has finished")

	logger.Info("Database migration has started")
	if err := migrations.Up(cfg, logger); err != nil {
		logger.Error("failed to apply migrations", slog.String("error", err.Error()))
		os.Exit(1)
	}
	logger.Info("Database migration has finished")

	pgDb, err := db.InitDB(cfg)
	if err != nil {
		logger.Error("failed to initialize database", slog.String("error", err.Error()))
		os.Exit(1)
	}

	repo := repository.InitRepository(pgDb)
	routerBuilder := router.InitBuilder(repo)
	routerBuilder.UserRouters()

	logger.Info("The server has started up at :8081")
	if err := routerBuilder.GetEngine().Run(":8081"); err != nil {
		logger.Error("failed to start the server", slog.String("error", err.Error()))
		os.Exit(1)
	}

	logger.Info("The server is running successfully")
}

func loggerInit() *slog.Logger {
	loggerHandler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(loggerHandler)
	slog.SetDefault(logger)
	return logger
}
