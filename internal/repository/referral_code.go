package repository

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ref_system/internal"
	"ref_system/internal/domain"
	"ref_system/internal/models"
	"ref_system/pkg/referral_code"
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
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, internal.ErrRecordNoFound
		}
		return nil, result.Error
	}

	return referralCodes, nil
}

func (r *ReferralCodeRepository) Create(referralCode *models.ReferralCode) (*domain.ReferralCode, error) {
	if referralCode.Code == "" {
		generatedCode, err := referral_code.GenerateUniqueReferralCode(r.repository.db)
		if err != nil {
			return nil, err
		}
		referralCode.Code = generatedCode
	}

	result := r.repository.db.Table("referral_codes").Create(&referralCode)
	if result.Error != nil {
		return nil, result.Error
	}

	newReferralCode := &domain.ReferralCode{
		UUID:      referralCode.UUID,
		OwnerUUID: referralCode.OwnerUUID,
		Code:      referralCode.Code,
		CreatedAt: referralCode.CreatedAt,
		UpdatedAt: referralCode.UpdatedAt,
		DeletedAt: gorm.DeletedAt{},
	}
	return newReferralCode, nil

}
func (r *ReferralCodeRepository) GetByCode(code string) (*domain.ReferralCode, error) {
	referralCode := &domain.ReferralCode{}
	result := r.repository.db.Table("referral_codes").Where("code = ?", code).First(&referralCode)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, internal.ErrRecordNoFound
		}
		return nil, result.Error
	}
	return referralCode, nil
}

func (r *ReferralCodeRepository) GetByOwnerUUID(uuid uuid.UUID) (*domain.ReferralCode, error) {
	var referralCode *domain.ReferralCode
	var user *domain.User
	userCheckResult := r.repository.db.Table("users").Where("uuid = ?", uuid).First(&user)
	if userCheckResult.Error != nil {
		if errors.Is(userCheckResult.Error, gorm.ErrRecordNotFound) {
			return nil, internal.ErrOwnerNotFound

		}
		return nil, userCheckResult.Error
	}

	result := r.repository.db.Table("referral_codes").Where("owner_uuid = ?", uuid).First(&referralCode)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, internal.ErrRecordNoFound
		}
		return nil, result.Error
	}
	return referralCode, nil
}

func (r *ReferralCodeRepository) CreateNewCodeAndNewUser(referralCode *models.ReferralCode) (*domain.ReferralCode, error) {
	if referralCode.Code == "" {
		generatedCode, err := referral_code.GenerateUniqueReferralCode(r.repository.db)
		if err != nil {
			return nil, internal.ErrCodeGenerate
		}
		referralCode.Code = generatedCode
	}

	tx := r.repository.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var user *domain.User
	if err := tx.Table("users").Where("uuid = ?", referralCode.OwnerUUID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user.UUID = referralCode.OwnerUUID
			if err := tx.Table("users").Create(&user).Error; err != nil {
				tx.Rollback()
				return nil, err
			}
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Table("referral_codes").Create(&referralCode).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	newReferralCode := &domain.ReferralCode{
		UUID:      referralCode.UUID,
		OwnerUUID: referralCode.OwnerUUID,
		Code:      referralCode.Code,
		CreatedAt: referralCode.CreatedAt,
		UpdatedAt: referralCode.UpdatedAt,
		DeletedAt: gorm.DeletedAt{},
	}
	return newReferralCode, nil
}
