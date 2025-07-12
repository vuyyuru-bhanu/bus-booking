package main

import (
	"bus-booking/config"
	"bus-booking/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDB() // ✅ Fix this line
	routes.SetupRoutes(r)

	r.Run(":8080")
}
