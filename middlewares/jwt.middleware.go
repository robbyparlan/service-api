package middlewares

import (
	"net/http"
	"service-api/config"
	"service-api/utils"

	echoJwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

/* JwtMiddleware provides a reusable JWT middleware */
func JwtMiddleware(cfg *config.Config) echo.MiddlewareFunc {
	secretKey := []byte(cfg.JwtSecretKey)

	config := echoJwt.Config{
		SigningKey: secretKey,
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, utils.CustomResponse{Status: http.StatusUnauthorized, Message: "Unauthorized. Token invalid", Data: nil})
		},
	}

	return echoJwt.WithConfig(config)
}
