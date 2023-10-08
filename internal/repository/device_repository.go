package repository

import (
	"errors"

	"github.com/frangar97/testapi/internal/entities"
	"gorm.io/gorm"
)

type DeviceRepository interface {
	GetAllDevices() ([]entities.Device, error)
	CreateDevice(*entities.Device) error
	GetDeviceById(string) (entities.Device, error)
	UpdateDevice(*entities.Device) error
	DeleteDevice(*entities.Device) error
}

type deviceRepositoryImpl struct {
	db *gorm.DB
}

func (u deviceRepositoryImpl) GetAllDevices() ([]entities.Device, error) {
	devices := []entities.Device{}
	result := u.db.Find(&devices)

	if result.Error != nil {
		return devices, result.Error
	}

	return devices, nil
}

func (u deviceRepositoryImpl) CreateDevice(device *entities.Device) error {
	result := u.db.Create(device)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("device was not created")
	}

	return nil
}

func (u deviceRepositoryImpl) GetDeviceById(username string) (entities.Device, error) {
	device := entities.Device{}

	result := u.db.First(&device, "id=?", username)
	if result.Error != nil {
		return device, result.Error
	}

	return device, nil
}

func (u deviceRepositoryImpl) UpdateDevice(device *entities.Device) error {
	result := u.db.Save(device)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("device was not updated")
	}

	return nil
}

func (u deviceRepositoryImpl) DeleteDevice(device *entities.Device) error {
	result := u.db.Delete(device)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("device was not deleted")
	}

	return nil
}
