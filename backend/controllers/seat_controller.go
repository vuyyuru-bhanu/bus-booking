package controllers

import (
	"net/http"
	"strconv"

	"bus-booking/config"
	"bus-booking/models"
	"github.com/gin-gonic/gin"
)

// GetAvailableSeats returns available seats for a bus
func GetAvailableSeats(c *gin.Context) {
	busIDStr := c.Param("bus_id")
	busID, err := strconv.Atoi(busIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bus ID"})
		return
	}

	seats, err := models.GetAvailableSeats(config.DB, busID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get available seats"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"seats": seats})
}

// SelectSeats allows user to select seats for a booking
func SelectSeats(c *gin.Context) {
	busIDStr := c.Param("bus_id")
	busID, err := strconv.Atoi(busIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bus ID"})
		return
	}

	var input struct {
		Seats []string `json:"seats"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	seatNumbers := []int{}
	for _, seatStr := range input.Seats {
		seatNum, err := strconv.Atoi(seatStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid seat number"})
			return
		}
		seatNumbers = append(seatNumbers, seatNum)
	}

	err = models.SelectSeats(config.DB, busID, seatNumbers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to select seats"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Seats selected successfully"})
}
