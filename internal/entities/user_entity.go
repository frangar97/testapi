package entities

import (
	"time"
)

type User struct {
	ID        string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserName  string `gorm:"not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
