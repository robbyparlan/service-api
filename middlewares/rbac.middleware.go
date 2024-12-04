package middlewares

import (
	"fmt"
	"net/http"
	"service-api/config"
	"service-api/utils"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// RbacMiddleware validates the user's role from JWT token.
func RbacMiddleware(cfg *config.Config, requiredRoles []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing or invalid token")
			}

			// Extract the token part from Authorization header
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing bearer token")
			}

			// Parse and validate the JWT token
			token, err := jwt.ParseWithClaims(tokenString, &utils.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				// Replace with your secret key
				return []byte(cfg.JwtSecretKey), nil
			})

			if err != nil || !token.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired token")
			}

			// Extract custom claims
			claims, ok := token.Claims.(*utils.JwtCustomClaims)
			if !ok || claims.Data.Role == "" {
				return echo.NewHTTPError(http.StatusForbidden, "missing role in token")
			}

			// Check if user's role is allowed
			for _, role := range requiredRoles {
				if claims.Data.Role == role {
					return next(c)
				}
			}

			// User's role not authorized
			return &utils.CustomError{
				StatusCode: http.StatusForbidden,
				Message:    fmt.Sprintf("User role %s not authorized", claims.Data.Role),
			}
		}
	}
}
