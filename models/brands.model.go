package models

import (
	"gorm.io/gorm"
)

type Brands struct {
	gorm.Model
	BrandName string `gorm:"size:225;not null" json:"BrandName"`
}
