package controllers

import (
	"log"
	"net/http"
	dtos "service-api/dtos"
	dtosBrand "service-api/dtos/brand"
	"service-api/services"
	"service-api/utils"

	"github.com/labstack/echo/v4"
)

type BrandController struct {
	brandService services.BrandService
}

func NewBrandController(brandService services.BrandService) *BrandController {
	return &BrandController{brandService: brandService}
}

func (c *BrandController) AddBrand(ctx echo.Context) error {
	var req dtosBrand.CreateBrandDTO

	if err := ctx.Bind(&req); err != nil {
		log.Println("Error binding request:", err.Error())
		return ctx.JSON(http.StatusBadRequest, utils.CustomResponse{Status: http.StatusBadRequest, Message: err.Error(), Data: nil})
	}

	//validate
	if err := ctx.Validate(req); err != nil {
		log.Println("Error validating request:", err.Error())
		validationErrors := utils.HandleValidationError(err)
		return ctx.JSON(http.StatusBadRequest, utils.CustomResponse{
			Status:  http.StatusBadRequest,
			Message: utils.MESSAGE_VALIDATION_ERROR,
			Data:    validationErrors,
		})
	}

	data, err := c.brandService.AddBrand(&req)
	if err != nil {
		log.Println("Error adding brand:", err.Error())
		customError := utils.NewCustomError(err)
		return ctx.JSON(customError.StatusCode, utils.CustomResponse{Status: customError.StatusCode, Message: customError.Message, Data: nil})
	}

	return ctx.JSON(http.StatusOK, utils.CustomResponse{Status: http.StatusOK, Message: utils.MESSAGE_SUCCESS, Data: data})
}

func (c *BrandController) ListBrand(ctx echo.Context) error {
	var req dtos.BaseRequestDTO

	if err := ctx.Bind(&req); err != nil {
		log.Println("Error binding request:", err.Error())
		return ctx.JSON(http.StatusBadRequest, utils.CustomResponse{Status: http.StatusBadRequest, Message: err.Error(), Data: nil})
	}

	//validate
	if err := ctx.Validate(req); err != nil {
		log.Println("Error validating request:", err.Error())
		validationErrors := utils.HandleValidationError(err)
		return ctx.JSON(http.StatusBadRequest, utils.CustomResponse{
			Status:  http.StatusBadRequest,
			Message: utils.MESSAGE_VALIDATION_ERROR,
			Data:    validationErrors,
		})
	}

	data, err := c.brandService.ListBrand(&req)
	if err != nil {
		log.Println("Error list brand:", err.Error())
		customError := utils.NewCustomError(err)
		return ctx.JSON(customError.StatusCode, utils.CustomResponse{Status: customError.StatusCode, Message: customError.Message, Data: nil})
	}

	return ctx.JSON(http.StatusOK, utils.CustomResponse{Status: http.StatusOK, Message: utils.MESSAGE_SUCCESS, Data: data})
}

func (c *BrandController) UpdateBrand(ctx echo.Context) error {
	var req dtosBrand.UpdateBrandDTO

	if err := ctx.Bind(&req); err != nil {
		log.Println("Error binding request:", err.Error())
		return ctx.JSON(http.StatusBadRequest, utils.CustomResponse{Status: http.StatusBadRequest, Message: err.Error(), Data: nil})
	}

	//validate
	if err := ctx.Validate(req); err != nil {
		log.Println("Error validating request:", err.Error())
		validationErrors := utils.HandleValidationError(err)
		return ctx.JSON(http.StatusBadRequest, utils.CustomResponse{
			Status:  http.StatusBadRequest,
			Message: utils.MESSAGE_VALIDATION_ERROR,
			Data:    validationErrors,
		})
	}

	data, err := c.brandService.UpdateBrand(&req)
	if err != nil {
		log.Println("Error update brand:", err.Error())
		customError := utils.NewCustomError(err)
		return ctx.JSON(customError.StatusCode, utils.CustomResponse{Status: customError.StatusCode, Message: customError.Message, Data: nil})
	}

	return ctx.JSON(http.StatusOK, utils.CustomResponse{Status: http.StatusOK, Message: utils.MESSAGE_SUCCESS, Data: data})
}

func (c *BrandController) DeleteBrand(ctx echo.Context) error {
	var req dtosBrand.DeleteBrandDTO

	if err := ctx.Bind(&req); err != nil {
		log.Println("Error binding request:", err.Error())
		return ctx.JSON(http.StatusBadRequest, utils.CustomResponse{Status: http.StatusBadRequest, Message: err.Error(), Data: nil})
	}

	//validate
	if err := ctx.Validate(req); err != nil {
		log.Println("Error validating request:", err.Error())
		validationErrors := utils.HandleValidationError(err)
		return ctx.JSON(http.StatusBadRequest, utils.CustomResponse{
			Status:  http.StatusBadRequest,
			Message: utils.MESSAGE_VALIDATION_ERROR,
			Data:    validationErrors,
		})
	}

	err := c.brandService.DeleteBrand(req.ID)
	if err != nil {
		log.Println("Error delete brand:", err.Error())
		customError := utils.NewCustomError(err)
		return ctx.JSON(customError.StatusCode, utils.CustomResponse{Status: customError.StatusCode, Message: customError.Message, Data: nil})
	}

	return ctx.JSON(http.StatusOK, utils.CustomResponse{Status: http.StatusOK, Message: utils.MESSAGE_SUCCESS, Data: nil})
}
