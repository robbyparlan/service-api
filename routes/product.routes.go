package routes

import (
	"service-api/config"
	"service-api/controllers"
	"service-api/middlewares"

	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

func RegisterProductRoutes(e *echo.Group, productController *controllers.ProductController, db *gorm.DB, cfg *config.Config) {
	r := e.Group("/product")
	r.Use(middlewares.JwtMiddleware(cfg))
	r.Use(middlewares.RbacMiddleware(cfg, []string{"admin", "user"}))

	// List Product Routes
	r.POST("/list", productController.ListProduct)
	r.POST("", productController.AddProduct)
	r.PUT("", productController.UpdateProduct)
	r.DELETE("", productController.DeleteProduct)
}
