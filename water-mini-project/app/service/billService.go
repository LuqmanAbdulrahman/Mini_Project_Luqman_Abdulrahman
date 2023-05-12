package service

import (
	"water-mini-project/app/handler"
	"github.com/labstack/echo/v4"
)

func GetBillByIdService(c echo.Context) error {
	return handler.GetBillByIdHandler(c)
}

func CreateBillService(c echo.Context) error {
	return handler.CreateBillHandler(c)
}

func UpdateBillService(c echo.Context) error {
	return handler.UpdateBillHandler(c)
}

func DeleteBillService(c echo.Context) error {
	return handler.DeleteBillHandler(c)
}