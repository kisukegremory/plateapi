package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kisukegremory/plateapi/internal/auth"
	"github.com/kisukegremory/plateapi/internal/broker"
	"github.com/kisukegremory/plateapi/internal/db"
)

func init() {
	db.ConnectToDB()
	db.SyncDatabase()
	broker.ConnectToBroker()
	broker.ConnectToChannel()
	broker.SyncMessageBroker()
}

func main() {
	router := gin.Default()
	router.GET("/v1/auth", AuthRoute)
	router.GET("/v1/auth/validate", auth.ValidationMiddleware, ValidateAuthRoute)
	router.GET("/v1/vehicles/:plate", auth.ValidationMiddleware, PlateRoute)
	router.Run()

	defer broker.ChannelConnection.Close()
	defer broker.BrokerConnection.Close()
}
