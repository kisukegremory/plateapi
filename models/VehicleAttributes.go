package models

import (
	"time"

	"gorm.io/gorm"
)

type VehicleAttributes struct {
	gorm.Model
	Plate        string `gorm:"PrimaryKey"`
	Year         uint32
	ModelYear    uint32
	Manufacturer string
	VehicleModel string
	SubModel     string
	Version      string
	Uf           string
	City         string
	Color        string
	Seats        uint8
	Fuel         string
	Created      time.Time `gorm:"index"`
}
