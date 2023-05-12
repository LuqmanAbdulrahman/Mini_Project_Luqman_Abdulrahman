package repository

import (
	"water-mini-project/app/service"
	"github.com/labstack/echo/v4"
)

func GetBillByIdRepository(c echo.Context) error {
	return service.GetBillByIdService(c)
}

func CreateBillRepository(c echo.Context) error {
	return service.CreateBillService(c)
}

func UpdateBillRepository(c echo.Context) error {
	return service.UpdateBillService(c)
}

func DeleteBillRepository(c echo.Context) error {
	return service.DeleteBillService(c)
}