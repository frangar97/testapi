package entities

import (
	"time"
)

type Firmware struct {
	ID           string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name         string    `gorm:"not null" json:"name"`
	DeviceID     string    `gorm:"type:uuid;not null;" json:"device_id"`
	Version      string    `json:"version"`
	ReleaseNotes string    `json:"release_notes"`
	ReleaseDate  time.Time `json:"release_date"`
	Url          string    `json:"url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
