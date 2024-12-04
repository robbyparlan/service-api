package models

import (
	"time"

	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	ProductName string  `gorm:"size:255;column:product_name;not null" json:"ProductName"`
	Price       float64 `gorm:"column:price;not null" json:"Price"`
	Quantity    int     `gorm:"column:quantity;not null" json:"Quantity"`
	BrandID     uint    `gorm:"column:brand_id;not null" json:"BrandID"`
}

type ProductWithBrand struct {
	ID          int       `json:"ID"`
	CreatedAt   time.Time `json:"CreatedAt"`
	UpdatedAt   time.Time `json:"UpdatedAt"`
	ProductName string    `json:"ProductName"`
	Price       float64   `json:"Price"`
	Quantity    int       `json:"Quantity"`
	BrandID     uint      `json:"BrandID"`
	BrandName   string    `json:"BrandName"`
}
