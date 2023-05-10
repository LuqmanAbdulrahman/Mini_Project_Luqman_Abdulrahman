package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewErrorResponse(code int, message string) *ErrorResponse {
	return &ErrorResponse{Code: code, Message: message}
}

func ErrorHandler(err error, c echo.Context) {
	if err != nil {
		var code int
		var message string

		switch e := err.(type) {
		case *echo.HTTPError:
			code = e.Code
			message = e.Message.(string)
		case *ErrorResponse:
			code = e.Code
			message = e.Message
		default:
			code = http.StatusInternalServerError
			message = "internal server error"
		}

		c.JSON(code, NewErrorResponse(code, message))
	}
}

// meng-handle error yang terjadi pada aplikasi dengan mengirimkan
//response yang sesuai berdasarkan jenis error yang terjadi
