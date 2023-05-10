package middleware

import (
	"net/http"

	"Mini_Project/internal/app/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, utils.NewErrorResponse(http.StatusUnauthorized, "missing authorization token"))
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusBadRequest, "invalid token")
			}
			return []byte("your-secret-key"), nil // replace with your own secret key
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, utils.NewErrorResponse(http.StatusUnauthorized, "invalid authorization token"))
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := claims["user_id"].(string)
			c.Set("userID", userID)
			return next(c)
		}

		return c.JSON(http.StatusUnauthorized, utils.NewErrorResponse(http.StatusUnauthorized, "invalid authorization token"))
	}
}

func ErrorHandler(err error, c echo.Context) {
	if err != nil {
		code := http.StatusInternalServerError
		message := "internal server error"

		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
			message = he.Message.(string)
		}

		c.JSON(code, utils.NewErrorResponse(code, message))
	}
}
