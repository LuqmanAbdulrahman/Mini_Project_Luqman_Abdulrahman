package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"Mini_Project/internal/app/model"
	"Mini_Project/internal/app/service"
	"Mini_Project/internal/app/utils"
)

type MeteranHandler struct {
	meteranService *service.MeteranService
}

func NewMeteranHandler(db *gorm.DB) *MeteranHandler {
	return &MeteranHandler{meteranService: service.NewMeteranService(db)}
}

func (h *MeteranHandler) CreateMeteran(c echo.Context) error {
	req := new(model.MeteranRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewErrorResponse(http.StatusBadRequest, "invalid request payload"))
	}

	userID := uuid.FromStringOrNil(c.Get("userID").(string))
	meteran, err := h.meteranService.CreateMeteran(req, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return c.JSON(http.StatusCreated, meteran)
}

func (h *MeteranHandler) GetMeterans(c echo.Context) error {
	meterans, err := h.meteranService.GetMeterans()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return c.JSON(http.StatusOK, meterans)
}

func (h *MeteranHandler) GetMeteranByID(c echo.Context) error {
	meteranID := uuid.FromStringOrNil(c.Param("id"))
	meteran, err := h.meteranService.GetMeteranByID(meteranID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return c.JSON(http.StatusOK, meteran)
}

func (h *MeteranHandler) UpdateMeteran(c echo.Context) error {
	req := new(model.MeteranRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewErrorResponse(http.StatusBadRequest, "invalid request payload"))
	}

	meteranID := uuid.FromStringOrNil(c.Param("id"))
	meteran, err := h.meteranService.UpdateMeteran(meteranID, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return c.JSON(http.StatusOK, meteran)
}

func (h *MeteranHandler) DeleteMeteran(c echo.Context) error {
	meteranID := uuid.FromStringOrNil(c.Param("id"))
	if err := h.meteranService.DeleteMeteran(meteranID); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "meteran deleted successfully",
	})
}
