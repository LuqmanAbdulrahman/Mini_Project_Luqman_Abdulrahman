package repository

import (
	"water-mini-project/app/service"
	"github.com/labstack/echo/v4"
)

func SearchCustomerRepository(c echo.Context) error {
	return service.SearchCustomerService(c)
}

func SearchBillRepository(c echo.Context) error {
	return service.SearchBillService(c)
}

func GetBillPerCustomerRepository(c echo.Context) error {
	return service.GetBillPerCustomerService(c)
}

func SearchByPeriodRepository(c echo.Context) error {
	return service.SearchByPeriodService(c)
}