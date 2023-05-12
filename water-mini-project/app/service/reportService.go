package service

import (
	"water-mini-project/app/handler"
	"github.com/labstack/echo/v4"
)

func SearchCustomerService(c echo.Context) error {
	return handler.SearchCustomerHandler(c)
}

func SearchBillService(c echo.Context) error {
	return handler.SearchBillHandler(c)
}

func GetBillPerCustomerService(c echo.Context) error {
	return handler.GetBillPerCustomerHandler(c)
}

func SearchByPeriodService(c echo.Context) error {
	return handler.SearchByPeriodHandler(c)
}