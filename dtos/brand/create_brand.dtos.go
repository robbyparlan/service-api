package dtos

type CreateBrandDTO struct {
	BrandName string `json:"BrandName" validate:"required,lte=255"`
}
