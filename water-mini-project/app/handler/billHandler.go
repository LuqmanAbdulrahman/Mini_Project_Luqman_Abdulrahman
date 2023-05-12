package handler

import (
	"net/http"
	"water-mini-project/app/models"
	"water-mini-project/app/utils"

	"github.com/labstack/echo/v4"
)

func GetBillByIdHandler(c echo.Context) error {
	id := c.Param("id")

	p := new(models.Bill)
	if err := utils.GetDB().Preload("Pelanggan").First(p, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Bill not found")
	}

	return c.JSON(http.StatusOK, p)
}

func CreateBillHandler(c echo.Context) error {
	p := new(models.Bill)

	if err := c.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := utils.GetDB().Create(p).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, p)
}

func UpdateBillHandler(c echo.Context) error {
	id := c.Param("id")

	p := new(models.Bill)
	if err := utils.GetDB().First(p, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Bill not found")
	}

	if err := c.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := utils.GetDB().Save(p).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, p)
}

func DeleteBillHandler(c echo.Context) error {
	id := c.Param("id")

	if err := utils.GetDB().Delete(&models.Bill{}, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
