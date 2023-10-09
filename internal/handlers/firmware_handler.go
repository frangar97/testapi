package handlers

import (
	"net/http"

	"github.com/frangar97/testapi/internal/models"
	"github.com/labstack/echo/v4"
)

func (h Handlers) CreateFirmwareHandler(c echo.Context) error {
	firmwareModel := models.FirmwareModel{}
	err := c.Bind(&firmwareModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, requestResponse{Message: err.Error()})
	}

	errores := firmwareModel.ValidateFirmware()
	if len(errores) != 0 {
		return c.JSON(http.StatusBadRequest, requestResponse{Data: errores})
	}

	firmware, err := h.services.FirmwareService.CreateFirmware(firmwareModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, requestResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, requestResponse{Data: firmware})
}

func (h Handlers) GetAllFirmwareHandler(c echo.Context) error {
	firmwares, err := h.services.FirmwareService.GetAllFirmwares()
	if err != nil {
		return c.JSON(http.StatusBadRequest, requestResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, requestResponse{Data: firmwares})
}

func (h Handlers) GetFirmwareByIdHandler(c echo.Context) error {
	firmwareId := c.Param("id")
	firmware, err := h.services.FirmwareService.GetFirmwareById(firmwareId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, requestResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, requestResponse{Data: firmware})
}

func (h Handlers) DeleteFirmwareByIdHandler(c echo.Context) error {
	firmwareId := c.Param("id")
	err := h.services.DeviceService.DeleteDevice(firmwareId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, requestResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, requestResponse{Message: "firmware deleted successfully"})
}

func (h Handlers) UpdateFirmwareHandler(c echo.Context) error {
	firmwareModel := models.FirmwareModel{}
	err := c.Bind(&firmwareModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, requestResponse{Message: err.Error()})
	}

	errores := firmwareModel.ValidateFirmware()
	if len(errores) != 0 {
		return c.JSON(http.StatusBadRequest, requestResponse{Data: errores})
	}

	firmwareId := c.Param("id")
	err = h.services.FirmwareService.UpdateFirmware(firmwareId, firmwareModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, requestResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, requestResponse{Message: "device updated successfully"})
}
