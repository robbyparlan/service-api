package routes

import (
	"service-api/config"
	"service-api/container"
	"service-api/utils"

	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config) {
	c := container.NewContainer(db)
	r := e.Group(utils.API_VERSION)

	/*
		Register routes
	*/
	RegisterAuthRoutes(r, c.AuthContainer.AuthController, db)
	RegisterBrandRoutes(r, c.BrandContainer.BrandController, db, cfg)
	RegisterProductRoutes(r, c.ProductContainer.ProductController, db, cfg)
}
