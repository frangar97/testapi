package repository

import (
	"errors"

	"github.com/frangar97/testapi/internal/entities"
	"gorm.io/gorm"
)

type FirmwareRepository interface {
	GetAllFirmwares() ([]entities.Firmware, error)
	CreateFirmware(*entities.Firmware) error
	GetFirmwareById(string) (entities.Firmware, error)
	UpdateFirmware(*entities.Firmware) error
	DeleteFirmware(*entities.Firmware) error
}

type firmwareRepositoryImpl struct {
	db *gorm.DB
}

func (u firmwareRepositoryImpl) GetAllFirmwares() ([]entities.Firmware, error) {
	firmwares := []entities.Firmware{}
	result := u.db.Find(&firmwares)

	if result.Error != nil {
		return firmwares, result.Error
	}

	return firmwares, nil
}

func (u firmwareRepositoryImpl) CreateFirmware(firmware *entities.Firmware) error {
	result := u.db.Create(firmware)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("firmware was not created")
	}

	return nil
}

func (u firmwareRepositoryImpl) GetFirmwareById(username string) (entities.Firmware, error) {
	firmware := entities.Firmware{}

	result := u.db.First(&firmware, "id=?", username)
	if result.Error != nil {
		return firmware, result.Error
	}

	return firmware, nil
}

func (u firmwareRepositoryImpl) UpdateFirmware(firmware *entities.Firmware) error {
	result := u.db.Save(firmware)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("firmware was not updated")
	}

	return nil
}

func (u firmwareRepositoryImpl) DeleteFirmware(firmware *entities.Firmware) error {
	result := u.db.Delete(firmware)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("firmware was not deleted")
	}

	return nil
}
