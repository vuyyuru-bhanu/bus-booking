package controllers

import (
	"net/http"

	"bus-booking/config"
	"bus-booking/models"
	"bus-booking/utils"
	"github.com/gin-gonic/gin"
)

// GetProfile returns the profile of the authenticated user
func GetProfile(c *gin.Context) {
	userID := utils.GetUserIDFromContext(c)
	user, err := models.GetUserByID(config.DB, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// UpdateProfile updates the profile of the authenticated user
func UpdateProfile(c *gin.Context) {
	userID := utils.GetUserIDFromContext(c)
	var input struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	err := models.UpdateUserProfile(config.DB, userID, input.Name, input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}
