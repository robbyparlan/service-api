package container

import (
	"service-api/controllers"
	"service-api/repository"
	"service-api/services"

	"gorm.io/gorm"
)

type AuthContainer struct {
	AuthController *controllers.AuthController
}

func NewAuthContainer(db *gorm.DB) *AuthContainer {
	authRepo := repository.NewAuthRepository(db)
	authService := services.NewAuthService(authRepo, db)
	authController := controllers.NewAuthController(authService)

	return &AuthContainer{AuthController: authController}
}
