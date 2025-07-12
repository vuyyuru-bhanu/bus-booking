package controllers

import (
	"net/http"

	"bus-booking/config"
	"bus-booking/models"
	"github.com/gin-gonic/gin"
)

// GetNotifications handles GET /notifications
func GetNotifications(c *gin.Context) {
	var notifications []models.Notification
	if err := config.DB.Order("created_at desc").Find(&notifications).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notifications"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"notifications": notifications})
}

// CreateNotification handles POST /api/admin/notifications
func CreateNotification(c *gin.Context) {
	var input struct {
		Message string `json:"message" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	notification := models.Notification{Message: input.Message}
	if err := config.DB.Create(&notification).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create notification"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"notification": notification})
}
