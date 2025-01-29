package routes

import (
	"service-api/config"
	"service-api/controllers"

	// "service-api/middlewares"

	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

func RegisterBrandRoutes(e *echo.Group, brandController *controllers.BrandController, db *gorm.DB, cfg *config.Config) {
	r := e.Group("/brand")
	// r.Use(middlewares.JwtMiddleware(cfg))
	// r.Use(middlewares.RbacMiddleware(cfg, []string{"admin"}))

	// List Brand Routes
	r.POST("/list", brandController.ListBrand)
	r.POST("", brandController.AddBrand)
	r.PUT("", brandController.UpdateBrand)
	r.DELETE("", brandController.DeleteBrand)
}
