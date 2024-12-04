package utils

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {

	if err := cv.Validator.Struct(i); err != nil {
		return err
	}

	return nil
}

/*
HandleValidationError provides a reusable HTTP error handler
*/
func HandleValidationError(err error) []ValidationErrorResponse {
	log.Printf("Raw validation error: %T - %v", err, err)
	var errors []ValidationErrorResponse

	// Pastikan error bisa dikonversi ke validator.ValidationErrors
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		// Jika tidak, kembalikan error generik
		return []ValidationErrorResponse{
			{Field: "unknown", Message: "Invalid validation error format"},
		}
	}

	// Loop through each error
	for _, err := range validationErrors {
		element := ValidationErrorResponse{
			Field:   err.Field(),
			Message: fmt.Sprintf("The %s field is %s", err.Field(), err.Tag()),
		}
		switch err.Tag() {
		case "required":
			element.Message = fmt.Sprintf("The %s field is required", err.Field())
		case "email":
			element.Message = fmt.Sprintf("The %s field is not a valid email address", err.Field())
		case "gte":
			element.Message = fmt.Sprintf("The %s field must be at least %s characters long", err.Field(), err.Param())
		case "lte":
			element.Message = fmt.Sprintf("The %s field must be at most %s characters long", err.Field(), err.Param())
		case "min":
			element.Message = fmt.Sprintf("The %s field must be at least %s", err.Field(), err.Param())
		case "max":
			element.Message = fmt.Sprintf("The %s field must be at most %s", err.Field(), err.Param())
		default:
			element.Message = fmt.Sprintf("The %s field is not valid", err.Field())
		}
		errors = append(errors, element)
	}
	return errors
}
