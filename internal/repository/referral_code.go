package repository

import (
	"ref_system/internal/domain"
)

type ReferralCodeRepository struct {
	repository *Repository
}

func NewReferralCodeRepository(repo *Repository) *ReferralCodeRepository {
	return &ReferralCodeRepository{
		repository: repo,
	}
}

func (r *ReferralCodeRepository) GetAll() ([]*domain.ReferralCode, error) {
	var referralCodes []*domain.ReferralCode
	result := r.repository.db.Table("referral_codes").Find(&referralCodes)
	if result.Error != nil {
		return nil, result.Error
	}
	return referralCodes, nil
}

//TODO: referral_code generator....
