package service

import (
	"water-mini-project/app/handler"
	"github.com/labstack/echo/v4"
)

func GetPelangganByIdService(c echo.Context) error {
	return handler.GetPelangganByIdHandler(c)
}

func CreatePelangganService(c echo.Context) error {
	return handler.CreatePelangganHandler(c)
}

func UpdatePelangganService(c echo.Context) error {
	return handler.UpdatePelangganHandler(c)
}

func DeletePelangganService(c echo.Context) error {
	return handler.DeletePelangganHandler(c)
}