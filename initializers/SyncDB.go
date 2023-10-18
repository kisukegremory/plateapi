package initializers

import "github.com/kisukegremory/plateapi/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.VehiclePlates{})
	DB.AutoMigrate(&models.VehicleAttributes{})
}
