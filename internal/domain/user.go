package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	CreatedAt time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt time.Time      `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" db:"deleted_at"`
	UUID      *uuid.UUID     `json:"uuid" validate:"required,uuid"`
}
