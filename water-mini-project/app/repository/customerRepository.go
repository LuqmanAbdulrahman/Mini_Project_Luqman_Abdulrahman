package repository

import (
	"water-mini-project/app/service"
	"github.com/labstack/echo/v4"
)

func GetPelangganByIdRepository(c echo.Context) error {
	return service.GetPelangganByIdService(c)
}

func CreatePelangganRepository(c echo.Context) error {
	return service.CreatePelangganService(c)
}

func UpdatePelangganRepository(c echo.Context) error {
	return service.UpdatePelangganService(c)
}

func DeletePelangganRepository(c echo.Context) error {
	return service.DeletePelangganService(c)
}