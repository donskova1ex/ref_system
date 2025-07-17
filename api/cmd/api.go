package main

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"os"
	"ref_system/internal/config"
	"ref_system/migrations"
)

func main() {
	loggerHandler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(loggerHandler)
	slog.SetDefault(logger)

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

	logger.Info("The server has started up")
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err = router.Run(":8080")
	if err != nil {
		logger.Error("failed to start server", slog.String("error", err.Error()))
	}
	logger.Info("The server is running successfully")
}
