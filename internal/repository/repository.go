package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository     UserRepository
	DeviceRepository   DeviceRepository
	FirmwareRepository FirmwareRepository
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		UserRepository:     userRepositoryImpl{db: db},
		DeviceRepository:   deviceRepositoryImpl{db: db},
		FirmwareRepository: firmwareRepositoryImpl{db: db},
	}
}
