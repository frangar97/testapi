package handlers

import (
	"net/http"

	"github.com/frangar97/testapi/internal/models"
	"github.com/labstack/echo/v4"
)

func (h Handlers) CreateDeviceHandler(c echo.Context) error {
	deviceModel := models.DeviceModel{}
	err := c.Bind(&deviceModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, requestResponse{Message: err.Error()})
	}

	errores := deviceModel.ValidateDevice()
	if len(errores) != 0 {
		return c.JSON(http.StatusBadRequest, requestResponse{Data: errores})
	}

	device, err := h.services.DeviceService.CreateDevice(deviceModel.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, requestResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, requestResponse{Data: device})
}

func (h Handlers) GetAllDevicesHandler(c echo.Context) error {
	devices, err := h.services.DeviceService.GetAllDevices()
	if err != nil {
		return c.JSON(http.StatusBadRequest, requestResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, requestResponse{Data: devices})
}

func (h Handlers) GetDeviceByIdHandler(c echo.Context) error {
	deviceId := c.Param("id")
	device, err := h.services.DeviceService.GetDeviceById(deviceId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, requestResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, requestResponse{Data: device})
}

func (h Handlers) DeleteDeviceByIdHandler(c echo.Context) error {
	deviceId := c.Param("id")
	err := h.services.DeviceService.DeleteDevice(deviceId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, requestResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, requestResponse{Message: "device deleted successfully"})
}

func (h Handlers) UpdateDeviceHandler(c echo.Context) error {
	deviceModel := models.DeviceModel{}
	err := c.Bind(&deviceModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, requestResponse{Message: err.Error()})
	}

	errores := deviceModel.ValidateDevice()
	if len(errores) != 0 {
		return c.JSON(http.StatusBadRequest, requestResponse{Data: errores})
	}

	deviceId := c.Param("id")
	err = h.services.DeviceService.UpdateDevice(deviceId, deviceModel.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, requestResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, requestResponse{Message: "device updated successfully"})
}
