package repository

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ref_system/internal"
	"ref_system/internal/domain"
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

func (u *UserRepository) Create(user *models.User) (*domain.User, error) {
	newUser := &domain.User{
		UUID: user.UUID,
	}
	result := u.repository.db.Table("users").Create(newUser)
	if result.Error != nil {
		return nil, result.Error
	}

	return newUser, nil
}
func (u *UserRepository) GetAll() ([]*domain.User, error) {
	var users []*domain.User
	result := u.repository.db.Table("users").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
func (u *UserRepository) GetByUUID(uuid *uuid.UUID) (*domain.User, error) {
	var user *domain.User
	result := u.repository.db.Table("users").Where("uuid = ?", uuid).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, internal.ErrRecordNoFound
		}
		return nil, result.Error
	}

	return user, nil
}
