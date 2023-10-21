package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kisukegremory/plateapi/internal/auth"
	"github.com/kisukegremory/plateapi/internal/initializers"
	"github.com/kisukegremory/plateapi/internal/models"
	plate "github.com/kisukegremory/plateapi/internal/plate"
)

func PlateRoute(c *gin.Context) {
	plate_string := c.Param("plate")
	match, _ := plate.PlateValidate(plate_string)
	if !match {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Wrong Plate"})
	}

	vehicleRequest := models.VehiclePlates{
		ID:      uuid.NewString(), // in the future we change to uuid4
		UserId:  uuid.NewString(),
		Plate:   plate_string,
		Created: time.Now(),
	}

	err := initializers.SendMessage(vehicleRequest)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Problems on publishing the message"})
	}
	c.String(http.StatusAccepted, ("Sent to queue: " + plate_string))
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
