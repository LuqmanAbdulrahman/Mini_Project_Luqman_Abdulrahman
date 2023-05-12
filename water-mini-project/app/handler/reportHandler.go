package handler

import (
	"fmt"
	"net/http"
	"time"
	"water-mini-project/app/models"
	"water-mini-project/app/utils"

	"github.com/labstack/echo/v4"
)

func SearchCustomerHandler(c echo.Context) error {
	query := c.Param("query")
	p := new(models.Customer)
	if err:= utils.GetDB().Preload("Laporan").Where(
		"nama LIKE ? OR email LIKE ? OR umur LIKE ?",
		"%"+query+"%",
		"%"+query+"%",
		"%"+query+"%",
	).Find(p).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Data not found")
	}
	return c.JSON(http.StatusOK, p)
}

func SearchBillHandler(c echo.Context) error {
	query := c.Param("query")
	fmt.Println(query)
	p := new(models.Bill)
	if err:= utils.GetDB().Preload("Pelanggan").Where(
		"total_debit LIKE ? OR harga_air LIKE ? OR pelanggan_id LIKE ? OR created_at LIKE ?",
		"%"+query+"%",
		"%"+query+"%",
		"%"+query+"%",
		"%"+query+"%",
	).Find(p).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Data not found")
	}
	return c.JSON(http.StatusOK, p)
}

func GetBillPerCustomerHandler(c echo.Context) error {
	return GetPelangganByIdHandler(c)
}

type SearchRequest struct {
    StartDate time.Time `json:"start_date"`
    EndDate   time.Time `json:"end_date"`
}

// EXAMPLE BODY
// {
// 	"start_date": "2023-01-01T00:00:00Z",
// 	"end_date": "2023-12-31T23:59:59Z"
// }

func SearchByPeriodHandler(c echo.Context) error {
	searchRequest := new(SearchRequest)
    if err := c.Bind(&searchRequest); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }

	p := new(models.Bill)
    err := utils.GetDB().Preload("Pelanggan").Where("created_at BETWEEN ? AND ? OR updated_at BETWEEN ? AND ?", searchRequest.StartDate, searchRequest.EndDate, searchRequest.StartDate, searchRequest.EndDate).Find(&p).Error
    if err != nil {
        return c.JSON(http.StatusInternalServerError, "Internal server error")
    }

    return c.JSON(http.StatusOK, p)
}
