package repository

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ref_system/internal"
	"ref_system/internal/models"
)

type UserRepository struct {
	repository *Repository
}

func NewUserRepository(repo *Repository) *UserRepository {
	return &UserRepository{
		repository: repo,
	}
}

func (u *UserRepository) Create(user *models.User) (*models.User, error) {
	result := u.repository.db.Table("users").Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
func (u *UserRepository) GetAll() ([]*models.User, error) {
	var users []*models.User
	result := u.repository.db.Table("users").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
func (u *UserRepository) GetByUUID(uuid *uuid.UUID) (*models.User, error) {
	var user *models.User
	result := u.repository.db.Table("users").Where("uuid = ?", uuid).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, internal.ErrRecordNoFound
		}
		return nil, result.Error
	}

	return user, nil
}
