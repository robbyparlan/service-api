package services

import (
	"errors"
	"net/http"
	dtos "service-api/dtos"
	dtosBrand "service-api/dtos/brand"
	"service-api/models"
	"service-api/repository"
	"service-api/utils"
	"strings"

	"gorm.io/gorm"
)

type BrandService interface {
	AddBrand(product *dtosBrand.CreateBrandDTO) (*models.Brands, error)
	ListBrand(baseRequest *dtos.BaseRequestDTO) (*dtos.BaseResponseDTO[models.Brands], error)
	UpdateBrand(product *dtosBrand.UpdateBrandDTO) (*models.Brands, error)
	DeleteBrand(id int) error
}

type brandService struct {
	repo repository.BrandRepository
	db   *gorm.DB
}

func NewBrandService(repo repository.BrandRepository, db *gorm.DB) BrandService {
	return &brandService{repo: repo, db: db}
}

func (s *brandService) AddBrand(product *dtosBrand.CreateBrandDTO) (*models.Brands, error) {
	// Check if product name already exists
	data, err := s.repo.CheckBrand(s.db, strings.TrimSpace(product.BrandName))
	if err != nil && err.Error() != "record not found" {
		return nil, err
	}

	if data.ID > 0 {
		return nil, errors.New("brand name already exists")
	}

	return s.repo.AddBrand(&models.Brands{
		BrandName: product.BrandName,
	})
}

func (s *brandService) ListBrand(baseRequest *dtos.BaseRequestDTO) (*dtos.BaseResponseDTO[models.Brands], error) {
	data, totalRecord, page, pageSize, totalPages, err := s.repo.ListBrand(s.db, baseRequest.Page, baseRequest.PageSize, baseRequest.Search)

	if err != nil {
		return nil, err
	}

	return &dtos.BaseResponseDTO[models.Brands]{
		Values:      *data,
		CurrentPage: page,
		PageSize:    pageSize,
		TotalItems:  totalRecord,
		TotalPages:  totalPages,
	}, nil
}

func (s *brandService) UpdateBrand(brand *dtosBrand.UpdateBrandDTO) (*models.Brands, error) {
	// Check if brand exists before updating it
	checkBrand, err := s.repo.FindBrand(s.db, brand.ID)
	if err != nil {
		return nil, err
	}

	if checkBrand == nil {
		return nil, &utils.CustomError{StatusCode: http.StatusNotFound, Message: "Brand not found", Err: nil}
	}

	checkBrand.BrandName = strings.TrimSpace(brand.BrandName)

	return s.repo.UpdateBrand(s.db, checkBrand)
}

func (s *brandService) DeleteBrand(id int) error {
	// Check if brand exists before deleting it
	checkBrand, err := s.repo.FindBrand(s.db, id)
	if err != nil {
		return err
	}

	if checkBrand == nil {
		return &utils.CustomError{StatusCode: http.StatusNotFound, Message: "Brand not found", Err: nil}
	}

	return s.repo.DeleteBrand(s.db, id)
}
