package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"bus-booking/config"
	"bus-booking/models"
)

type BookingInput struct {
	BusID      uint `json:"bus_id" binding:"required"`
	SeatNumber int  `json:"seat_number" binding:"required"`
}

// CreateBooking handles POST /bookings to create a new booking for a user.
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

	err := config.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&booking).Error; err != nil {
			return err
		}

		bus.AvailableSeats--
		if err := tx.Save(&bus).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"booking": booking})
}

// CancelBooking cancels a booking by ID
func CancelBooking(c *gin.Context) {
	id := c.Param("id")
	var booking models.Booking
	if err := config.DB.First(&booking, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	booking.Status = "cancelled"
	if err := config.DB.Save(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel booking"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking cancelled successfully"})
}
