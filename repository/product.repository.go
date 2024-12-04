package repository

import (
	"log"
	"service-api/models"
	"service-api/utils"

	"gorm.io/gorm"
)

type ProductRepository interface {
	AddProduct(product *models.Products) (*models.Products, error)
	ListProduct(db *gorm.DB, page int, pageSize int, search string) (*[]models.ProductWithBrand, int64, int64, int64, int64, error)
	UpdateProduct(db *gorm.DB, product *models.Products) (*models.Products, error)
	DeleteProduct(db *gorm.DB, id int) error
	FindProduct(db *gorm.DB, id int) (*models.Products, error)
	CheckProduct(db *gorm.DB, productName string, brandId int) (*models.Products, error)
	FindBrand(db *gorm.DB, id int) (*models.Brands, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) AddProduct(product *models.Products) (*models.Products, error) {
	return product, r.db.Create(product).Error
}

func (r *productRepository) ListProduct(db *gorm.DB, page int, pageSize int, search string) (*[]models.ProductWithBrand, int64, int64, int64, int64, error) {
	// var product []models.Products
	var productWithBrand []models.ProductWithBrand
	var totalRecords int64
	productField := utils.GetSelectFields(models.Products{}, true)
	productField = append(productField, "brands.brand_name")
	productField = append(productField, "products.id", "products.created_at", "products.updated_at")

	log.Println("productField", productField)

	// Hitung total record untuk pagination
	if err := db.Model(&models.Products{}).Joins("LEFT JOIN brands ON products.brand_id = brands.id").Count(&totalRecords).Error; err != nil {
		return nil, 0, 0, 0, 0, err
	}

	// logic pagination
	offset := (page - 1) * pageSize
	query := db.Model(&models.Products{}).
		Select(productField).
		Joins("LEFT JOIN brands ON products.brand_id = brands.id")

	if search != "" {
		query = query.Where("(products.product_name iLIKE ? OR brands.brand_name iLIKE ?)", "%"+search+"%", "%"+search+"%")
	}

	err := query.
		Offset(offset).
		Limit(pageSize).
		Find(&productWithBrand).Error

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

	return &productWithBrand, totalRecords, int64(page), int64(pageSize), totalPages, nil
}

func (r *productRepository) UpdateProduct(db *gorm.DB, product *models.Products) (*models.Products, error) {
	return product, r.db.Save(product).Error
}

func (r *productRepository) DeleteProduct(db *gorm.DB, id int) error {
	return r.db.Where("id = ?", id).Delete(&models.Products{}).Error
}

func (r *productRepository) FindProduct(db *gorm.DB, id int) (*models.Products, error) {
	var product *models.Products
	return product, r.db.First(&product, id).Error
}

func (r *productRepository) CheckProduct(db *gorm.DB, productName string, brandId int) (*models.Products, error) {
	var product *models.Products
	return product, r.db.Where("product_name = ?", productName).Where("brand_id = ?", brandId).First(&product).Error
}

func (r *productRepository) FindBrand(db *gorm.DB, id int) (*models.Brands, error) {
	var brand *models.Brands
	return brand, r.db.First(&brand, id).Error
}
