package repository

import (
	"net/http"
	"service-api/models"
	"service-api/utils"

	"gorm.io/gorm"
)

type BrandRepository interface {
	AddBrand(product *models.Brands) (*models.Brands, error)
	ListBrand(db *gorm.DB, page int, pageSize int, search string) (*[]models.Brands, int64, int64, int64, int64, error)
	UpdateBrand(db *gorm.DB, product *models.Brands) (*models.Brands, error)
	DeleteBrand(db *gorm.DB, id int) error
	FindBrand(db *gorm.DB, id int) (*models.Brands, error)
	CheckBrand(db *gorm.DB, brandName string) (*models.Brands, error)
}

type brandRepository struct {
	db *gorm.DB
}

func NewBrandRepository(db *gorm.DB) BrandRepository {
	return &brandRepository{db: db}
}

func (r *brandRepository) AddBrand(brand *models.Brands) (*models.Brands, error) {
	return brand, r.db.Create(brand).Error
}

func (r *brandRepository) ListBrand(db *gorm.DB, page int, pageSize int, search string) (*[]models.Brands, int64, int64, int64, int64, error) {
	var brand []models.Brands
	var totalRecords int64
	brandField := utils.GetSelectFields(models.Brands{}, true)

	// Hitung total record untuk pagination
	if err := db.Model(&models.Brands{}).Count(&totalRecords).Error; err != nil {
		return nil, 0, 0, 0, 0, err
	}

	// logic pagination
	offset := (page - 1) * pageSize
	query := db.Model(&models.Brands{}).
		Select(brandField)

	if search != "" {
		query = query.Where("(brands.brand_name iLIKE ?)", "%"+search+"%")
	}

	err := query.
		Offset(offset).
		Limit(pageSize).
		Find(&brand).Error

	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	// Hitung total page
	totalPages := int64(0)
	if totalRecords > 0 {
		totalPages = totalRecords / int64(pageSize)
		if totalRecords%int64(pageSize) != 0 {
			totalPages++
		}
	}

	return &brand, totalRecords, int64(page), int64(pageSize), totalPages, nil
}

func (r *brandRepository) UpdateBrand(db *gorm.DB, brand *models.Brands) (*models.Brands, error) {
	return brand, r.db.Save(brand).Error
}

func (r *brandRepository) DeleteBrand(db *gorm.DB, id int) error {
	// Check if used in product
	var totalProduct int64
	err := db.Model(&models.Products{}).Where("brand_id = ?", id).Count(&totalProduct).Error
	if err != nil {
		return err
	}

	if totalProduct > 0 {
		return &utils.CustomError{StatusCode: http.StatusBadRequest, Message: "Brand is used in product", Err: nil}
	}

	return r.db.Where("id = ?", id).Delete(&models.Brands{}).Error
}

func (r *brandRepository) FindBrand(db *gorm.DB, id int) (*models.Brands, error) {
	var brand *models.Brands
	return brand, r.db.First(&brand, id).Error
}

func (r *brandRepository) CheckBrand(db *gorm.DB, brandName string) (*models.Brands, error) {
	var brand *models.Brands
	return brand, r.db.Where("brand_name = ?", brandName).First(&brand).Error
}
