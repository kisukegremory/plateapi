package models

import (
	"time"
)

type VehiclePlates struct {
	ID      string    `gorm:"PrimaryKey" json:"id"`
	UserId  string    `json:"user_id"`
	Plate   string    `gorm:"index" json:"plate"`
	Created time.Time `gorm:"index" json:"created"`
}
