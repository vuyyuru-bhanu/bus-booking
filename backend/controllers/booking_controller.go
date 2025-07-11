package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"bus-booking/config"
	"bus-booking/models"
)

type BookingInput struct {
	BusID      uint `json:"bus_id" binding:"required"`
	SeatNumber int  `json:"seat_number" binding:"required"`
}

func CreateBooking(c *gin.Context) {
	var input BookingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := uint(c.MustGet("user_id").(float64))

	var bus models.Bus
	if err := config.DB.First(&bus, input.BusID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bus not found"})
		return
	}

	if bus.AvailableSeats < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No available seats"})
		return
	}

	booking := models.Booking{
		UserID:     userID,
		BusID:      input.BusID,
		SeatNumber: input.SeatNumber,
		Status:     "confirmed",
	}

	if err := config.DB.Create(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		return
	}

	bus.AvailableSeats--
	if err := config.DB.Save(&bus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update bus seats"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"booking": booking})
}