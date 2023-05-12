package repository

import (
	"water-mini-project/app/service"

	"github.com/labstack/echo/v4"
)

func GetAuthGenerateTokenRepository(c echo.Context) error {
	return service.GetAuthGenerateTokenService(c)
}

func ValidateAuthTokenRepository(next echo.HandlerFunc) echo.HandlerFunc {
	return service.ValidateAuthTokenService(next)
}