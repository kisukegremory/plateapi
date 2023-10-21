package broker

import (
	"encoding/json"
	"fmt"

	"github.com/kisukegremory/plateapi/internal/models"
)

func PublishPlate(vehicle models.VehiclePlates) error {

	msgBody, err := json.Marshal(vehicle)

	if err != nil {
		return fmt.Errorf("problems on parsing the json: %v", err)
	}

	return PublishJson(SearchQueue.Name, msgBody)
}

func PublishStore(vehicle models.Vehicle) error {

	msgBody, err := json.Marshal(vehicle)

	if err != nil {
		return fmt.Errorf("problems on parsing the json: %v", err)
	}

	return PublishJson(StoreQueue.Name, msgBody)
}
