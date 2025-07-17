package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Referral struct {
	gorm.Model
	UUID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	ReferrerID uint      `gorm:"not null"`
	Referrer   User      `gorm:"foreignKey:ReferrerID"`
	RefereeID  uint      `gorm:"not null;uniqueIndex"`
	Referee    User      `gorm:"foreignKey:RefereeID"`
	CodeID     uint
	Code       ReferralCode `gorm:"foreignKey:CodeID"`
}
