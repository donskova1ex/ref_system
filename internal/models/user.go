package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex" validate:"omitempty,uuid4"`
	UserContactInfo string    `gorm:"uniqueIndex" validate:"omitempty"`
}
