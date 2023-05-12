package handler

import (
	"net/http"
	"water-mini-project/app/models"
	"water-mini-project/app/utils"

	"github.com/labstack/echo/v4"
)

func GetPelangganByIdHandler(c echo.Context) error {
	id := c.Param("id")

	p := new(models.Customer)
	if err := utils.GetDB().Preload("Laporan").First(p, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Pelanggan not found")
	}

	return c.JSON(http.StatusOK, p)
}

func CreatePelangganHandler(c echo.Context) error {
	p := new(models.Customer)

	if err := c.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := utils.GetDB().Create(p).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, p)
}

func UpdatePelangganHandler(c echo.Context) error {
	id := c.Param("id")

	p := new(models.Customer)
	if err := utils.GetDB().First(p, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Pelanggan not found")
	}

	if err := c.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := utils.GetDB().Save(p).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, p)
}

func DeletePelangganHandler(c echo.Context) error {
	id := c.Param("id")

	if err := utils.GetDB().Delete(&models.Customer{}, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
