package service

import (
	"errors"

	"github.com/frangar97/testapi/internal/entities"
	"github.com/frangar97/testapi/internal/repository"
	"gorm.io/gorm"
)

type DeviceService interface {
	CreateDevice(string) (entities.Device, error)
	GetDeviceById(string) (entities.Device, error)
	GetAllDevices(string) ([]entities.Device, error)
}

type deviceServiceImpl struct {
	secret           string
	deviceRepository repository.DeviceRepository
}

func (d deviceServiceImpl) CreateDevice(name string) (entities.Device, error) {
	device := entities.Device{Name: name}
	err := d.deviceRepository.CreateDevice(&device)

	if err != nil {
		return device, err
	}

	return device, nil
}

func (d deviceServiceImpl) GetDeviceById(id string) (entities.Device, error) {

	device, err := d.deviceRepository.GetDeviceById(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return device, errors.New("no device found with the requested id")
		} else {
			return device, err
		}
	}

	return device, nil
}

func (d deviceServiceImpl) GetAllDevices(id string) ([]entities.Device, error) {
	return d.deviceRepository.GetAllDevices()
}