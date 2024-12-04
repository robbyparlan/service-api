package services

import (
	"errors"
	"net/http"
	dtos "service-api/dtos"
	dtosProduct "service-api/dtos/product"
	"service-api/models"
	"service-api/repository"
	"service-api/utils"
	"strings"

	"gorm.io/gorm"
)

type ProductService interface {
	AddProduct(product *dtosProduct.CreateProductDTO) (*models.Products, error)
	ListProduct(baseRequest *dtos.BaseRequestDTO) (*dtos.BaseResponseDTO[models.ProductWithBrand], error)
	UpdateProduct(product *dtosProduct.UpdateProductDTO) (*models.Products, error)
	DeleteProduct(id int) error
}

type productService struct {
	repo repository.ProductRepository
	db   *gorm.DB
}

func NewProductService(repo repository.ProductRepository, db *gorm.DB) ProductService {
	return &productService{repo: repo, db: db}
}

func (s *productService) AddProduct(product *dtosProduct.CreateProductDTO) (*models.Products, error) {
	// Check if product name already exists
	data, err := s.repo.CheckProduct(s.db, strings.TrimSpace(product.ProductName), product.BrandID)
	if err != nil && err.Error() != "record not found" {
		return nil, err
	}

	if data.ID > 0 {
		return nil, errors.New("product name already exists")
	}

	return s.repo.AddProduct(&models.Products{
		ProductName: strings.TrimSpace(product.ProductName),
		Price:       product.Price,
		Quantity:    product.Quantity,
		BrandID:     uint(product.BrandID),
	})
}

func (s *productService) ListProduct(baseRequest *dtos.BaseRequestDTO) (*dtos.BaseResponseDTO[models.ProductWithBrand], error) {
	data, totalRecord, page, pageSize, totalPages, err := s.repo.ListProduct(s.db, baseRequest.Page, baseRequest.PageSize, baseRequest.Search)

	if err != nil {
		return nil, err
	}

	return &dtos.BaseResponseDTO[models.ProductWithBrand]{
		Values:      *data,
		CurrentPage: page,
		PageSize:    pageSize,
		TotalItems:  totalRecord,
		TotalPages:  totalPages,
	}, nil
}

func (s *productService) UpdateProduct(product *dtosProduct.UpdateProductDTO) (*models.Products, error) {
	// Check if product exists before updating it
	checkProduct, err := s.repo.FindProduct(s.db, product.ID)
	if err != nil {
		return nil, err
	}

	if checkProduct == nil {
		return nil, &utils.CustomError{StatusCode: http.StatusNotFound, Message: "Product not found", Err: nil}
	}

	checkProduct.ProductName = strings.TrimSpace(product.ProductName)
	checkProduct.Price = product.Price
	checkProduct.Quantity = product.Quantity
	checkProduct.BrandID = uint(product.BrandID)

	return s.repo.UpdateProduct(s.db, checkProduct)
}

func (s *productService) DeleteProduct(id int) error {
	// Check if product exists before deleting it
	checkProduct, err := s.repo.FindProduct(s.db, id)
	if err != nil {
		return err
	}

	if checkProduct == nil {
		return &utils.CustomError{StatusCode: http.StatusNotFound, Message: "Product not found", Err: nil}
	}

	return s.repo.DeleteProduct(s.db, id)
}
