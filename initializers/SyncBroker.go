package initializers

import (
	"context"
	"time"

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

func SendMessage() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ChannelConnection.PublishWithContext(
		ctx,
		"",
		Queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello World!"),
		},
	)

}
