package models

import (
	"time"

	"gorm.io/gorm"
)

type VehiclePlates struct {
	gorm.Model
	ID      string `gorm:"PrimaryKey"`
	UserId  string
	Plate   string    `gorm:"index"`
	Created time.Time `gorm:"index"`
}
