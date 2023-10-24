package main

import (
	"encoding/json"
	"log/slog"
	"time"

	"github.com/kisukegremory/plateapi/internal/apiplaca"
	"github.com/kisukegremory/plateapi/internal/broker"
	"github.com/kisukegremory/plateapi/internal/db"
	"github.com/kisukegremory/plateapi/internal/models"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm/clause"
)

func VehicleSearchConsumer(msg amqp.Delivery) {
	var err error
	slog.Info("Received a message: %s", msg.Body)
	var vehicleSimulation models.VehiclePlates
	err = json.Unmarshal(msg.Body, &vehicleSimulation)
	broker.FailOnError(err, "Problems on decoding the VehiclePlate")

	var vehicleAttributes apiplaca.VehicleAttributesAPI
	vehicleAttributes, err = apiplaca.GetVehicleAttributesByPlate(vehicleSimulation.Plate)
	broker.FailOnError(err, "Problems on Finding the Vehicle")
	slog.Info("Attributes found: %v", vehicleAttributes)

	vehicle := models.Vehicle{
		Plate: vehicleSimulation,
		Attributes: models.VehicleAttributes{
			Plate:        vehicleAttributes.Plate,
			Year:         vehicleAttributes.Year,
			ModelYear:    vehicleAttributes.ModelYear,
			Manufacturer: vehicleAttributes.Manufacturer,
			VehicleModel: vehicleAttributes.VehicleModel,
			SubModel:     vehicleAttributes.SubModel,
			Version:      vehicleAttributes.Version,
			Uf:           vehicleAttributes.Uf,
			City:         vehicleAttributes.City,
			Color:        vehicleAttributes.Color,
			Origin:       vehicleAttributes.Origin,
			Created:      time.Now(),
		},
	}

	broker.PublishStore(vehicle)
	broker.FailOnError(err, "Problems on Publishing the Vehicle Store Queue")

}

func VehicleStoreConsumer(msg amqp.Delivery) {
	var err error
	slog.Info("Received a message: %s", msg.Body)
	var vehicle models.Vehicle
	err = json.Unmarshal(msg.Body, &vehicle)
	broker.FailOnError(err, "Problems on decoding the Vehicle")

	db.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&vehicle.Plate)
	db.DB.Create(&vehicle.Attributes)

}
