package dtos

type UpdateBrandDTO struct {
	ID        int    `json:"ID" validate:"required"`
	BrandName string `json:"BrandName" validate:"required,gte=5,lte=255"`
}
