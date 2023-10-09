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
	GetAllDevices() ([]entities.Device, error)
	UpdateDevice(string, string) error
	DeleteDevice(string) error
}

type deviceServiceImpl struct {
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

func (d deviceServiceImpl) GetAllDevices() ([]entities.Device, error) {
	return d.deviceRepository.GetAllDevices()
}

func (d deviceServiceImpl) UpdateDevice(id, name string) error {
	device, err := d.deviceRepository.GetDeviceById(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("no device found with the requested id")
		} else {
			return err
		}
	}

	device.Name = name
	return d.deviceRepository.UpdateDevice(&device)
}

func (d deviceServiceImpl) DeleteDevice(id string) error {
	device, err := d.deviceRepository.GetDeviceById(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("no device found with the requested id")
		} else {
			return err
		}
	}

	return d.deviceRepository.DeleteDevice(&device)
}
