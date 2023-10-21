package main

import (
	"log"

	"github.com/kisukegremory/plateapi/internal/broker"
)

func init() {
	broker.ConnectToBroker()
	broker.ConnectToChannel()
	broker.SyncMessageBroker()
}

func main() {

	msgs, err := broker.ChannelConnection.Consume(
		broker.Queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	broker.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for msg := range msgs {
			VehicleSearchConsumer(msg)
		}
	}()

	log.Printf("* Waiting for messages")
	<-forever

	defer broker.ChannelConnection.Close()
	defer broker.BrokerConnection.Close()

}
