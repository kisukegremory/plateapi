package main

import (
	"log"

	"github.com/kisukegremory/plateapi/internal/broker"
	"github.com/kisukegremory/plateapi/internal/db"
)

func init() {
	broker.ConnectToBroker()
	broker.ConnectToChannel()
	broker.SyncMessageBroker()
	db.ConnectToDB()
	db.SyncDatabase()
}

func main() {

	searchMsgs, err := broker.ChannelConnection.Consume(
		broker.SearchQueue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	broker.FailOnError(err, "Failed to register search consumer")

	storeMsgs, err := broker.ChannelConnection.Consume(
		broker.StoreQueue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	broker.FailOnError(err, "Failed to register store consumer")

	var forever chan struct{}

	go func() {
		for msg := range searchMsgs {
			VehicleSearchConsumer(msg)
		}
	}()

	go func() {
		for msg := range storeMsgs {
			VehicleStoreConsumer(msg)
		}
	}()

	log.Printf("* Waiting for messages")
	<-forever

	defer broker.ChannelConnection.Close()
	defer broker.BrokerConnection.Close()

}
