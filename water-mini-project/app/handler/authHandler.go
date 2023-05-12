package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

func AuthGenerateTokenHandler(c echo.Context) error {
	signingKey := []byte("your-secret-key")

	claims := jwt.MapClaims{
		"username": "therealuser",
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, tokenString)
}

func ValidateTokenHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")

		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token not provided")
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			signingKey := []byte("your-secret-key")
			return signingKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token signature")
			}
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid token")
		}

		if token.Valid {
			return next(c)
		}

		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
	}
}