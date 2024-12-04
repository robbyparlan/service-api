package routes

import (
	"service-api/container"
	"service-api/utils"

	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	c := container.NewContainer(db)
	r := e.Group(utils.API_VERSION)

	/*
		Register routes
	*/
	RegisterProductRoutes(r, c.ProductContainer.ProductController, db)
}
