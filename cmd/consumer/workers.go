package main

import (
	"encoding/json"
	"log"

	"github.com/kisukegremory/plateapi/internal/apiplaca"
	"github.com/kisukegremory/plateapi/internal/broker"
	"github.com/kisukegremory/plateapi/internal/models"
	amqp "github.com/rabbitmq/amqp091-go"
)

func VehicleSearchConsumer(msg amqp.Delivery) {
	var err error
	log.Printf("Received a message: %s", msg.Body)
	var vehicleSimulation models.VehiclePlates
	err = json.Unmarshal(msg.Body, &vehicleSimulation)
	broker.FailOnError(err, "Problems on decoding the VehiclePlate")

	var vehicleAttributes apiplaca.VehicleAttributesAPI
	vehicleAttributes, err = apiplaca.GetVehicleAttributesByPlate(vehicleSimulation.Plate)
	broker.FailOnError(err, "Problems on Finding the Vehicle")
	log.Printf("Attributes found: %v", vehicleAttributes)

}
