package services

import (
	"net/http"
	"service-api/config"
	dtosAuth "service-api/dtos/auth"
	"service-api/models"
	"service-api/repository"
	"service-api/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(*dtosAuth.LoginAuthDTO) (*dtosAuth.LoginAuthResponseDTO[models.Users], error)
}

type authService struct {
	repo repository.AuthRepository
	db   *gorm.DB
}

func NewAuthService(repo repository.AuthRepository, db *gorm.DB) AuthService {
	return &authService{repo: repo, db: db}
}

func (s *authService) Login(user *dtosAuth.LoginAuthDTO) (*dtosAuth.LoginAuthResponseDTO[models.Users], error) {
	userMdl := &models.Users{
		Username: user.Username,
		Password: user.Password,
	}
	data, err := s.repo.Login(userMdl)
	if err != nil {
		return nil, err
	}

	if data.ID == 0 {
		return nil, &utils.CustomError{
			StatusCode: http.StatusNotFound,
			Message:    "User not found",
		}
	}

	hashPassword := data.Password
	data.Password = ""
	err = bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(user.Password))
	if err != nil {
		return nil, &utils.CustomError{
			StatusCode: http.StatusUnauthorized,
			Message:    "Invalid password",
		}
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}
	token, err := utils.GenerateJWT(int(data.ID), data, cfg)
	if err != nil {
		return nil, err
	}

	return &dtosAuth.LoginAuthResponseDTO[models.Users]{
		User:        *data,
		AccessToken: token,
	}, nil
}
