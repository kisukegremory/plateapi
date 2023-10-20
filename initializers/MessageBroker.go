package initializers

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var BrokerConnection *amqp.Connection
var ChannelConnection *amqp.Channel

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s : %s", msg, err)
	}
}

func ConnectToBroker() {
	var err error
	BrokerConnection, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Failed to Connect to Broker")
	log.Println("Sucessfully Connect to Broker")
}

func ConnectToChannel() {
	var err error
	ChannelConnection, err = BrokerConnection.Channel()
	FailOnError(err, "Failed to Connect to Broker Channel")
	log.Println("Sucessfully Connect to Broker Channel")
}
