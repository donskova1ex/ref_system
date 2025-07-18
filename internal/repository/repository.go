package repository

import (
	"gorm.io/gorm"
	"ref_system/pkg/db"
)

type Repository struct {
	db *gorm.DB
}

func InitRepository(db *db.DB) *Repository {
	return &Repository{
		db: db.DB,
	}
}
