package container

import (
	"service-api/controllers"
	"service-api/repository"
	"service-api/services"

	"gorm.io/gorm"
)

type BrandContainer struct {
	BrandController *controllers.BrandController
}

/*
Function NewTestContainer
Sebagai Dependency Injection repository service dan controller
*/
func NewBrandContainer(db *gorm.DB) *BrandContainer {
	brandRepo := repository.NewBrandRepository(db)                  // Mengembalikan interface
	brandService := services.NewBrandService(brandRepo, db)         // Mengembalikan interface
	brandController := controllers.NewBrandController(brandService) // Mengembalikan struct

	return &BrandContainer{
		BrandController: brandController,
	}
}
