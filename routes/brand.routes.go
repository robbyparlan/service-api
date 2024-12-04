package routes

import (
	"service-api/controllers"

	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

func RegisterBrandRoutes(e *echo.Group, brandController *controllers.BrandController, db *gorm.DB) {
	r := e.Group("/brand")
	r.POST("/list", brandController.ListBrand)
	r.POST("", brandController.AddBrand)
	r.PUT("", brandController.UpdateBrand)
	r.DELETE("", brandController.DeleteBrand)
}
