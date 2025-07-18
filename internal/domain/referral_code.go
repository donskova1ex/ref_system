package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ReferralCode struct {
	UUID      uuid.UUID      `json:"uuid" db:"uuid"`
	OwnerUUID uuid.UUID      `json:"owner_uuid" db:"owner_uuid"`
	Code      string         `json:"code" db:"code"`
	CreatedAt time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" db:"deleted_at"`
}
