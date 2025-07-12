package controllers

import (
	"net/http"
	"strconv"

	"bus-booking/config"
	"bus-booking/models"
	"github.com/gin-gonic/gin"
)

// GetReviews returns reviews for a bus
func GetReviews(c *gin.Context) {
	busIDStr := c.Param("bus_id")
	busID, err := strconv.Atoi(busIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bus ID"})
		return
	}

	reviews, err := models.GetReviewsByBusID(config.DB, busID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get reviews"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reviews": reviews})
}

// AddReview adds a review for a bus
func AddReview(c *gin.Context) {
	busIDStr := c.Param("bus_id")
	busID, err := strconv.Atoi(busIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bus ID"})
		return
	}

	var input struct {
		UserID  int    `json:"user_id"`
		Rating  int    `json:"rating"`
		Comment string `json:"comment"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	review := &models.Review{
		UserID:  uint(input.UserID),
		BusID:   uint(busID),
		Rating:  input.Rating,
		Comment: input.Comment,
	}

	err = models.AddReview(config.DB, review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add review"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review added successfully"})
}
