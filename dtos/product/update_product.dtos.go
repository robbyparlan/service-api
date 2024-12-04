package dtos

type UpdateProductDTO struct {
	ID          int     `json:"ID" validate:"required"`
	ProductName string  `json:"ProductName" validate:"required,gte=5,lte=255"`
	Price       float64 `json:"Price" validate:"required,gte=0,lte=1000000"`
	Quantity    int     `json:"Quantity" validate:"required,gte=0,lte=1000000"`
	BrandID     int     `json:"BrandID" validate:"required"`
}
