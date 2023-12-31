package broker

import (
	"log/slog"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

var SearchQueue *amqp.Queue
var StoreQueue *amqp.Queue

var BrokerConnection *amqp.Connection
var ChannelConnection *amqp.Channel

func ConnectToBroker() {
	var err error
	brokerDsn := os.Getenv("RABBIT_DSN")
	BrokerConnection, err = amqp.Dial(brokerDsn)
	FailOnError(err, "Failed to Connect to Broker")
	slog.Info("Sucessfully Connect to Broker")
}

func ConnectToChannel() {
	var err error
	ChannelConnection, err = BrokerConnection.Channel()
	FailOnError(err, "Failed to Connect to Broker Channel")
	slog.Info("Sucessfully Connect to Broker Channel")
}

func SyncMessageBroker() {
	var err error
	searchQueue, err := ChannelConnection.QueueDeclare(
		"vehicles.search",
		false,
		false,
		false,
		false,
		nil,
	)
	FailOnError(err, "Failed to Declare Search Queue")
	SearchQueue = &searchQueue

	storeQueue, err := ChannelConnection.QueueDeclare(
		"vehicles.store",
		false,
		false,
		false,
		false,
		nil,
	)
	FailOnError(err, "Failed to Declare Store Queue")
	StoreQueue = &storeQueue

}
