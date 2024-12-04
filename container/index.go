package container

import (
	"gorm.io/gorm"
)

type Container struct {
	ProductContainer *ProductContainer
}

func NewContainer(db *gorm.DB) *Container {

	return &Container{
		ProductContainer: NewProductContainer(db),
	}
}
