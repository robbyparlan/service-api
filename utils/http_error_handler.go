package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

/*
HTTPErrorHandler provides a reusable HTTP error handler
*/
func HTTPErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.Logger().Error(report)
	customError := NewCustomError(err)
	c.JSON(customError.StatusCode, CustomResponse{Status: customError.StatusCode, Message: customError.Message, Data: nil})
}
