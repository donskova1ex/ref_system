package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UUID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	UserID     uint      // ID пользователя
	User       User      `gorm:"foreignKey:UserID""`
	ReferralID *uint     // Опциональная ссылка на реферала
	Referral   *Referral `gorm:"foreignKey:ReferralID"`
	Amount     float64   `gorm:"type:decimal(10,2)"`
	Type       string    `gorm:"size:20;check:type IN ('credit','debit')"`
}
