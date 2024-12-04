package middlewares

import (
	"fmt"
	"time"

	"service-api/utils"

	"github.com/labstack/echo/v4"
)

func LoggingMiddleware(logger utils.LoggerInterface) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestID := c.Request().Header.Get("X-Request-ID")
			if requestID == "" {
				requestID = fmt.Sprintf("%d", time.Now().UnixNano()) // Default request ID if not provided
			}

			// Log request
			logger.LogRequest(c, requestID)

			// Wrap response writer
			rw := utils.NewResponseWriter(c.Response().Writer)
			c.Response().Writer = rw

			// Execute next middleware/handler
			err := next(c)

			// Log response
			logger.LogResponse(c, rw, requestID, err)
			return err
		}
	}
}
