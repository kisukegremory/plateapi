package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	plate "github.com/kisukegremory/plateapi/cmd/plate"
)

func PlateToQueue(c *gin.Context) {
	plate_string := c.Param("plate")
	match, _ := plate.PlateValidate(plate_string)
	switch match {
	case true:
		c.String(http.StatusAccepted, ("Sent to queue: " + plate_string))
	case false:
		c.String(http.StatusBadRequest, "Wrong Plate")
	}
}

func main() {
	router := gin.Default()
	router.GET("/v1/vehicles/:plate", PlateToQueue)
	router.Run()
}
