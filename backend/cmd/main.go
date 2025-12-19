package main

import (
	"bus-booking/config"
	"bus-booking/models"
	"bus-booking/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDB()
	
	// AutoMigrate the schema
	config.DB.AutoMigrate(&models.User{}, &models.Route{}, &models.Bus{}, &models.Booking{}, &models.BusImage{}, &models.Notification{}, &models.Seat{}, &models.Review{})

	routes.SetupRoutes(r)

	r.Run(":8080")
}
