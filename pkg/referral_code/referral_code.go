package referral_code

import (
	"gorm.io/gorm"
	"math/rand"
	"ref_system/internal"
	"ref_system/internal/models"
)

const (
	codeLength = 10
)

var (
	allowedChars = []rune("ABCDEFGHJKLMNPQRSTUVWXYZ123456789")
)

func generateReferralCode() string {
	code := make([]rune, codeLength)
	for i := range code {
		code[i] = allowedChars[rand.Intn(len(allowedChars))]
	}
	return string(code)
}

func GenerateUniqueReferralCode(db *gorm.DB) (string, error) {
	maxAttemps := 100
	for i := 0; i < maxAttemps; i++ {
		code := generateReferralCode()
		var count int64
		db.Model(&models.ReferralCode{}).Where("referral_code = ?", code).Count(&count)
		if count == 0 {
			return code, nil
		}
	}
	return "", internal.ErrCodeGenerate
}
