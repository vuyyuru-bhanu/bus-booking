package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"bus-booking/config"
	"bus-booking/models"
)

func GetRoutes(c *gin.Context) {
	var routes []models.Route
	if err := config.DB.Find(&routes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch routes"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"routes": routes})
}
