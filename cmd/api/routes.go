package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kisukegremory/plateapi/internal/auth"
	"github.com/kisukegremory/plateapi/internal/broker"
	"github.com/kisukegremory/plateapi/internal/db"
	"github.com/kisukegremory/plateapi/internal/models"
	plate "github.com/kisukegremory/plateapi/internal/plate"
)

func PlateRoute(c *gin.Context) {
	plate_string := c.Param("plate")
	match, _ := plate.PlateValidate(plate_string)
	if !match {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Wrong Plate"})
	}

	vehicleAttributes := models.VehicleAttributes{Plate: plate_string}

	result := db.DB.First(&vehicleAttributes)

	if result.Error == nil {
		vehicleAttributesJson, _ := json.Marshal(vehicleAttributes)
		c.AbortWithStatusJSON(200, string(vehicleAttributesJson))
		return
	}

	vehicleRequest := models.VehiclePlates{
		ID:      uuid.NewString(), // in the future we change to uuid4
		UserId:  uuid.NewString(),
		Plate:   plate_string,
		Created: time.Now(),
	}

	err := broker.PublishPlate(vehicleRequest)

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
