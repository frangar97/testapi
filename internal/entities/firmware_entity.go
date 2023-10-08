package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Firmware struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name         string    `gorm:"not null"`
	DeviceID     uuid.UUID `gorm:"type:uuid;not null;"`
	Version      string
	ReleaseNotes string
	ReleaseDate  time.Time
	Url          string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
