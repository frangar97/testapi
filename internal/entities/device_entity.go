package entities

import (
	"time"
)

type Device struct {
	ID        string     `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name      string     `gorm:"not null" json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Firmwares []Firmware `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
}
