package container

import (
	"service-api/controllers"
	"service-api/repository"
	"service-api/services"

	"gorm.io/gorm"
)

type ProductContainer struct {
	ProductController *controllers.ProductController
}

/*
Function NewTestContainer
Sebagai Dependency Injection repository service dan controller
*/
func NewProductContainer(db *gorm.DB) *ProductContainer {
	productRepo := repository.NewProductRepository(db)                    // Mengembalikan interface
	productService := services.NewProductService(productRepo, db)         // Mengembalikan interface
	productController := controllers.NewProductController(productService) // Mengembalikan struct

	return &ProductContainer{
		ProductController: productController,
	}
}
