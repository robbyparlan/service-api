package controllers

import (
	"log"
	"net/http"
	dtos "service-api/dtos"
	dtosProduct "service-api/dtos/product"
	"service-api/services"
	"service-api/utils"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (c *ProductController) AddProduct(ctx echo.Context) error {
	var req dtosProduct.CreateProductDTO

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

	data, err := c.productService.AddProduct(&req)
	if err != nil {
		log.Println("Error adding product:", err.Error())
		customError := utils.NewCustomError(err)
		return ctx.JSON(customError.StatusCode, utils.CustomResponse{Status: customError.StatusCode, Message: customError.Message, Data: nil})
	}

	return ctx.JSON(http.StatusOK, utils.CustomResponse{Status: http.StatusOK, Message: utils.MESSAGE_SUCCESS, Data: data})
}

func (c *ProductController) ListProduct(ctx echo.Context) error {
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

	data, err := c.productService.ListProduct(&req)
	if err != nil {
		log.Println("Error list product:", err.Error())
		customError := utils.NewCustomError(err)
		return ctx.JSON(customError.StatusCode, utils.CustomResponse{Status: customError.StatusCode, Message: customError.Message, Data: nil})
	}

	return ctx.JSON(http.StatusOK, utils.CustomResponse{Status: http.StatusOK, Message: utils.MESSAGE_SUCCESS, Data: data})
}

func (c *ProductController) UpdateProduct(ctx echo.Context) error {
	var req dtosProduct.UpdateProductDTO

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

	data, err := c.productService.UpdateProduct(&req)
	if err != nil {
		log.Println("Error update product:", err.Error())
		customError := utils.NewCustomError(err)
		return ctx.JSON(customError.StatusCode, utils.CustomResponse{Status: customError.StatusCode, Message: customError.Message, Data: nil})
	}

	return ctx.JSON(http.StatusOK, utils.CustomResponse{Status: http.StatusOK, Message: utils.MESSAGE_SUCCESS, Data: data})
}

func (c *ProductController) DeleteProduct(ctx echo.Context) error {
	var req dtosProduct.DeleteProductDTO

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

	err := c.productService.DeleteProduct(req.ID)
	if err != nil {
		log.Println("Error delete product:", err.Error())
		customError := utils.NewCustomError(err)
		return ctx.JSON(customError.StatusCode, utils.CustomResponse{Status: customError.StatusCode, Message: customError.Message, Data: nil})
	}

	return ctx.JSON(http.StatusOK, utils.CustomResponse{Status: http.StatusOK, Message: utils.MESSAGE_SUCCESS, Data: nil})
}
