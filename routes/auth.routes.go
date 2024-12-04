package routes

import (
	"service-api/controllers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterAuthRoutes(e *echo.Group, authController *controllers.AuthController, db *gorm.DB) {
	r := e.Group("/auth")
	r.POST("/login", authController.LoginAuthUser)
}
