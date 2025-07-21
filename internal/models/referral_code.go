package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReferralCode struct {
	gorm.Model
	UUID      uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4()"`
	OwnerUUID *uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`
	Owner     User       `gorm:"foreignKey:OwnerUUID;references:UUID"`
	Code      string     `gorm:"uniqueIndex;size:32;not null"`
}
