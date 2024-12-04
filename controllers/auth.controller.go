package controllers

import (
	"log"
	"net/http"
	dtosAuth "service-api/dtos/auth"
	"service-api/services"
	"service-api/utils"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) LoginAuthUser(ctx echo.Context) error {
	var req dtosAuth.LoginAuthDTO

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

	data, err := c.authService.Login(&req)
	if err != nil {
		log.Println("Error login:", err.Error())
		customError := utils.NewCustomError(err)
		return ctx.JSON(customError.StatusCode, utils.CustomResponse{Status: customError.StatusCode, Message: customError.Message, Data: nil})
	}

	return ctx.JSON(http.StatusOK, utils.CustomResponse{Status: http.StatusOK, Message: utils.MESSAGE_SUCCESS, Data: data})
}
