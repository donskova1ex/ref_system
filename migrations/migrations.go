package migrations

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
	"os"
	"ref_system/internal/models"
)

func Up(logger *slog.Logger) error {

	err := godotenv.Load(".env.local")
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
		os.Exit(1)
	}

	slog.Info("Configuration loaded")

	db, err := gorm.Open(postgres.Open(os.Getenv("DB_DSN")), &gorm.Config{})
	if err != nil {
		fmt.Errorf("failed to connect to database: %w", err)
		os.Exit(1)
	}
	logger.Info("Successfully connected to database")

	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
		return fmt.Errorf("failed creating extension \"uuid-ossp\": %w", err)
		os.Exit(1)
	}

	err = initMigrations(db)
	if err != nil {
		return fmt.Errorf("failed initialising migrations: %w", err)
	}

	return nil
}

func initMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		err := db.Migrator().DropTable(&models.User{})
		if err != nil {
			return fmt.Errorf("error dropping table: %w", err)
		}
		return fmt.Errorf("failed to create table Users: %w", err)

	}

	err = db.AutoMigrate(&models.ReferralCode{})
	if err != nil {
		err := db.Migrator().DropTable(&models.ReferralCode{})
		if err != nil {
			return fmt.Errorf("failed to drop table: %w", err)
		}
		return fmt.Errorf("failed to create table Referrals: %w", err)
	}

	err = db.AutoMigrate(&models.Referral{})
	if err != nil {
		err := db.Migrator().DropTable(&models.Referral{})
		if err != nil {
			return fmt.Errorf("failed to drop table Referrals: %w", err)
		}
		return fmt.Errorf("failed to create table Referrals: %w", err)
	}

	err = db.AutoMigrate(&models.Transaction{})
	if err != nil {
		err := db.Migrator().DropTable(&models.Transaction{})
		if err != nil {
			return fmt.Errorf("failed to drop table Transactions: %w", err)
		}
		return fmt.Errorf("failed to create table Transactions: %w", err)
	}
	return nil
}
