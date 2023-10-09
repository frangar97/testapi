package service

import (
	"errors"
	"time"

	"github.com/frangar97/testapi/internal/entities"
	"github.com/frangar97/testapi/internal/models"
	"github.com/frangar97/testapi/internal/repository"
	"gorm.io/gorm"
)

type FirmwareService interface {
	CreateFirmware(models.FirmwareModel) (entities.Firmware, error)
	UpdateFirmware(string, models.FirmwareModel) error
	GetFirmwareById(string) (entities.Firmware, error)
	GetAllFirmwares() ([]entities.Firmware, error)
	DeleteFirmware(string) error
}

type firmwareServiceImpl struct {
	firmwareRepository repository.FirmwareRepository
	deviceRepository   repository.DeviceRepository
}

func (f firmwareServiceImpl) GetAllFirmwares() ([]entities.Firmware, error) {
	return f.firmwareRepository.GetAllFirmwares()
}

func (f firmwareServiceImpl) DeleteFirmware(id string) error {
	firmware, err := f.firmwareRepository.GetFirmwareById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("no firmware found with the requested id")
		}

		return err
	}

	return f.firmwareRepository.DeleteFirmware(&firmware)
}

func (f firmwareServiceImpl) GetFirmwareById(id string) (entities.Firmware, error) {
	firmware, err := f.firmwareRepository.GetFirmwareById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return firmware, errors.New("no firmware found with the requested id")
		}

		return firmware, err
	}

	return firmware, nil
}

func (f firmwareServiceImpl) CreateFirmware(firmware models.FirmwareModel) (entities.Firmware, error) {
	date, err := time.Parse("2006-01-02", firmware.ReleaseDate)

	if err != nil {
		return entities.Firmware{}, err
	}

	_, err = f.deviceRepository.GetDeviceById(firmware.DeviceID)

	firmwareEntity := entities.Firmware{Name: firmware.Name, Version: firmware.Version, DeviceID: firmware.DeviceID, Url: firmware.Url, ReleaseNotes: firmware.ReleaseNotes, ReleaseDate: date}

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return firmwareEntity, errors.New("device not found with the selected deviceId")
		}

		return firmwareEntity, err
	}

	err = f.firmwareRepository.CreateFirmware(&firmwareEntity)

	if err != nil {
		return firmwareEntity, err
	}

	return firmwareEntity, nil
}

func (f firmwareServiceImpl) UpdateFirmware(id string, firmware models.FirmwareModel) error {
	firmwareEntity, err := f.firmwareRepository.GetFirmwareById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("no firmware found with the requested id")
		}

		return err
	}

	date, err := time.Parse("2006-01-02", firmware.ReleaseDate)

	if err != nil {
		return err
	}

	_, err = f.deviceRepository.GetDeviceById(firmware.DeviceID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("device not found with the selected deviceId")
		}

		return err
	}

	firmwareEntity.Name = firmware.Name
	firmwareEntity.Version = firmware.Version
	firmwareEntity.DeviceID = firmware.DeviceID
	firmwareEntity.ReleaseDate = date
	firmwareEntity.ReleaseNotes = firmware.ReleaseNotes
	firmwareEntity.Url = firmware.Url

	return f.firmwareRepository.UpdateFirmware(&firmwareEntity)
}
