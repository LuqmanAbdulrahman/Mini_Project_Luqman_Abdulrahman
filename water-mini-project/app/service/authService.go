package service

import (
	"water-mini-project/app/handler"

	"github.com/labstack/echo/v4"
)

func GetAuthGenerateTokenService(c echo.Context) error {
	return handler.AuthGenerateTokenHandler(c)
}

func ValidateAuthTokenService(next echo.HandlerFunc) echo.HandlerFunc {
	return handler.ValidateTokenHandler(next)
}