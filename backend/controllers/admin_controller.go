package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"github.com/gin-gonic/gin"
	"bus-booking/config"
	"bus-booking/models"
)

type RouteInput struct {
	Origin        string `json:"origin" binding:"required"`
	Destination   string `json:"destination" binding:"required"`
	DepartureTime string `json:"departure_time" binding:"required"`
	Duration      int    `json:"duration" binding:"required"` // Duration in minutes
}

type BusInput struct {
	RouteID       uint    `json:"route_id" binding:"required"`
	Company       string  `json:"company" binding:"required"`
	AC            bool    `json:"ac" binding:"required"`
	Type          string  `json:"type" binding:"required,oneof=seater sleeper"`
	Capacity      int     `json:"capacity" binding:"required"`
	AvailableSeats int    `json:"available_seats" binding:"required"`
	Price         float64 `json:"price" binding:"required"`
	Amenities     []string `json:"amenities"`
}

type NotificationInput struct {
	Message string `json:"message" binding:"required"`
}

func CreateRoute(c *gin.Context) {
	var input RouteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	route := models.Route{
		Origin:        input.Origin,
		Destination:   input.Destination,
		DepartureTime: input.DepartureTime,
		Duration:      input.Duration,
	}

	if err := config.DB.Create(&route).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create route"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"route": route})
}

func CreateBus(c *gin.Context) {
	var input BusInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bus := models.Bus{
		RouteID:       input.RouteID,
		Company:       input.Company,
		AC:            input.AC,
		Type:          input.Type,
		Capacity:      input.Capacity,
		AvailableSeats: input.AvailableSeats,
		Price:         input.Price,
		Amenities:     input.Amenities,
	}

	if err := config.DB.Create(&bus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bus"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"bus": bus})
}

func UploadBusImage(c *gin.Context) {
	busID := c.Param("bus_id")
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image file required"})
		return
	}

	// Validate file type and size (JPEG/PNG, max 5MB)
	if !isValidImage(file) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image format or size"})
		return
	}

	filename := fmt.Sprintf("%s-%s", busID, file.Filename)
	filePath := filepath.Join("uploads", filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	image := models.BusImage{
		BusID:    uint(c.MustGet("bus_id").(float64)),
		ImageURL: "/uploads/" + filename,
	}

	if err := config.DB.Create(&image).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"image": image})
}

func CreateNotification(c *gin.Context) {
	var input NotificationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	notification := models.Notification{
		Message: input.Message,
	}

	if err := config.DB.Create(&notification).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create notification"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"notification": notification})
}

func GetNotifications(c *gin.Context) {
	var notifications []models.Notification
	if err := config.DB.Find(&notifications).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notifications"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"notifications": notifications})
}

func GetAllBookings(c *gin.Context) {
	var bookings []models.Booking
	if err := config.DB.Find(&bookings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookings"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"bookings": bookings})
}

func isValidImage(file *multipart.FileHeader) bool {
	allowedExt := []string{".jpg", ".jpeg", ".png"}
	ext := strings.ToLower(filepath.Ext(file.Filename))
	for _, validExt := range allowedExt {
		if ext == validExt {
			return file.Size <= 5*1024*1024 // Max 5MB
		}
	}
	return false
}