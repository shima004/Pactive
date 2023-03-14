package model

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
	Name      string    `gorm:"unique"`
	Email     string
	Password  string
	PublicKey string
}
