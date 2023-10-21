package models

import (
	"time"
)

type VehicleAttributes struct {
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
	Origin       string
	Created      time.Time `gorm:"index"`
}
