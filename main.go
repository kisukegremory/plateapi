package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PlateToQueue(c *gin.Context) {
	plate := c.Param("plate")

	c.String(http.StatusAccepted, ("Sent to queue: " + plate))
}

func main() {
	router := gin.Default()
	router.GET("/v1/vehicles/:plate", PlateToQueue)
	router.Run()
}
