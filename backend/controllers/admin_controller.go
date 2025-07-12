package controllers

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"mime/multipart"

	"github.com/gin-gonic/gin"

	"bus-booking/config"
	"bus-booking/models"
)

type BusInput struct {
	RouteID        uint     `json:"route_id"`
	Company        string   `json:"company"`
	AC             bool     `json:"ac"`
	Type           string   `json:"type"`
	Capacity       int      `json:"capacity"`
	AvailableSeats int      `json:"available_seats"`
	Price          float64  `json:"price"`
	Amenities      []string `json:"amenities"`
}

func CreateBus(c *gin.Context) {
	var input BusInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	amenitiesJSON, err := json.Marshal(input.Amenities)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode amenities"})
		return
	}

	bus := models.Bus{
		RouteID:        input.RouteID,
		Company:        input.Company,
		AC:             input.AC,
		Type:           input.Type,
		Capacity:       input.Capacity,
		AvailableSeats: input.AvailableSeats,
		Price:          input.Price,
		Amenities:      string(amenitiesJSON),
	}

	if err := config.DB.Create(&bus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"bus": bus})
}

func isValidImage(file *multipart.FileHeader) bool {
	ext := strings.ToLower(filepath.Ext(file.Filename))
	switch ext {
	case ".jpg", ".jpeg", ".png":
		return true
	default:
		return false
	}
}

func CreateRoute(c *gin.Context) {
	var route models.Route
	if err := c.ShouldBindJSON(&route); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&route).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"route": route})
}

func UploadBusImage(c *gin.Context) {
	busIDParam := c.Param("busId")
	busID, err := strconv.ParseUint(busIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bus ID"})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	files := form.File["images"]
	var savedImages []models.BusImage

	for _, file := range files {
		if !isValidImage(file) {
			continue
		}
		filename := filepath.Base(file.Filename)
		savePath := filepath.Join("uploads", filename)

		if err := c.SaveUploadedFile(file, savePath); err != nil {
			continue
		}

		image := models.BusImage{
			BusID:    uint(busID),
			ImageURL: "/uploads/" + filename,
		}
		config.DB.Create(&image)
		savedImages = append(savedImages, image)
	}

	c.JSON(http.StatusOK, gin.H{"uploaded": savedImages})
}

func CreateNotification(c *gin.Context) {
	var notification models.Notification
	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&notification).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"notification": notification})
}

func GetAllBookings(c *gin.Context) {
	var bookings []models.Booking
	if err := config.DB.Preload("User").Preload("Bus").Find(&bookings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"bookings": bookings})
}

func GetNotifications(c *gin.Context) {
	var notifications []models.Notification

	if err := config.DB.Find(&notifications).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notifications"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"notifications": notifications})
}
