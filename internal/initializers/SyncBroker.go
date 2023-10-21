package initializers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/kisukegremory/plateapi/internal/models"
	amqp "github.com/rabbitmq/amqp091-go"
)

var Queue *amqp.Queue

func SyncMessageBroker() {
	var err error
	vehicleQueue, err := ChannelConnection.QueueDeclare(
		"vehicles.search",
		false,
		false,
		false,
		false,
		nil,
	)
	FailOnError(err, "Failed to Declare Vehicle Queue")
	Queue = &vehicleQueue

}

func SendMessage(vehicle models.VehiclePlates) error {

	msgBody, err := json.Marshal(vehicle)

	if err != nil {
		return fmt.Errorf("problems on parsing the json: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ChannelConnection.PublishWithContext(
		ctx,
		"",
		Queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msgBody,
		},
	)

	if err != nil {
		return fmt.Errorf("problems on publishing in the queue: %v", err)
	}

	return nil
}
