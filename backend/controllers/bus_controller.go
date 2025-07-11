package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"bus-booking/config"
	"bus-booking/models"
)

type BusSearchInput struct {
	Origin      string `form:"origin" binding:"required"`
	Destination string `form:"destination" binding:"required"`
	AC          *bool  `form:"ac"`
	Type        string `form:"type" binding:"omitempty,oneof=seater sleeper"`
	Company     string `form:"company"`
}

func GetBuses(c *gin.Context) {
	var input BusSearchInput
	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var buses []models.Bus
	query := config.DB.Preload("Route").Preload("Images").
		Joins("JOIN routes ON buses.route_id = routes.id").
		Where("routes.origin = ? AND routes.destination = ?", input.Origin, input.Destination)

	if input.AC != nil {
		query = query.Where("buses.ac = ?", *input.AC)
	}
	if input.Type != "" {
		query = query.Where("buses.type = ?", input.Type)
	}
	if input.Company != "" {
		query = query.Where("buses.company = ?", input.Company)
	}

	if err := query.Find(&buses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch buses"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"buses": buses})
}