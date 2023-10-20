package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kisukegremory/plateapi/auth"
	"github.com/kisukegremory/plateapi/initializers"
	plate "github.com/kisukegremory/plateapi/plate"
)

func PlateRoute(c *gin.Context) {
	plate_string := c.Param("plate")
	match, _ := plate.PlateValidate(plate_string)
	switch match {
	case true:
		c.String(http.StatusAccepted, ("Sent to queue: " + plate_string))
	case false:
		c.String(http.StatusBadRequest, "Wrong Plate")
	}
}

func AuthRoute(c *gin.Context) {
	token, err := auth.GenerateJwt()
	if err != nil {
		c.String(http.StatusUnauthorized, "Wrong Data")
		return
	}

	c.SetCookie("Authorization", token, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

func ValidateAuthRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ValidAuthenticated",
	})
}

func init() {
	initializers.ConnectToDB()
	initializers.SyncDatabase()
	initializers.ConnectToBroker()
	initializers.ConnectToChannel()
	initializers.SyncMessageBroker()
	initializers.SendMessage()
}

func main() {
	router := gin.Default()
	router.GET("/v1/auth", AuthRoute)
	router.GET("/v1/auth/validate", auth.ValidationMiddleware, ValidateAuthRoute)
	router.GET("/v1/vehicles/:plate", auth.ValidationMiddleware, PlateRoute)
	router.Run()

	defer initializers.ChannelConnection.Close()
	defer initializers.BrokerConnection.Close()
}
