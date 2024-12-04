package container

import (
	"gorm.io/gorm"
)

type Container struct {
	ProductContainer *ProductContainer
	BrandContainer   *BrandContainer
}

func NewContainer(db *gorm.DB) *Container {

	return &Container{
		ProductContainer: NewProductContainer(db),
		BrandContainer:   NewBrandContainer(db),
	}
}
