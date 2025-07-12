package controllers

import (
	"net/http"

	"bus-booking/config"
	"bus-booking/models"
	"bus-booking/utils"
	"github.com/gin-gonic/gin"
)

// GetBookingHistory returns the list of bookings for the authenticated user
func GetBookingHistory(c *gin.Context) {
	userID := utils.GetUserIDFromContext(c)
	bookings, err := models.GetBookingsByUserID(config.DB, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch booking history"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"bookings": bookings})
}

// GetBookingDetails returns details of a specific booking by ID for the authenticated user
func GetBookingDetails(c *gin.Context) {
	userID := utils.GetUserIDFromContext(c)
	bookingID := c.Param("id")

	booking, err := models.GetBookingByIDAndUserID(config.DB, bookingID, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"booking": booking})
}
