package controllers

import (
	"net/http"

	"bus-booking/config"
	"bus-booking/models"
	"github.com/gin-gonic/gin"
)

// CreateRoute handles POST /api/admin/routes
func CreateRoute(c *gin.Context) {
	var input models.Route
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create route"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"route": input})
}
