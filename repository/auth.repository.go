package repository

import (
	"service-api/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Login(user *models.Users) (*models.Users, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) Login(user *models.Users) (*models.Users, error) {
	return user, r.db.Where("username = ?", user.Username).First(user).Error
}
