package routes

import (
	"service-api/controllers"

	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

func RegisterProductRoutes(e *echo.Group, productController *controllers.ProductController, db *gorm.DB) {
	r := e.Group("/product")
	r.POST("/list", productController.ListProduct)
	r.POST("", productController.AddProduct)
	r.PUT("", productController.UpdateProduct)
	r.DELETE("", productController.DeleteProduct)
}
