FROM golang:1.21-alpine

WORKDIR /app

# Copy go.mod only (no go.sum yet)
COPY go.mod ./

# Initialize an empty go.sum by forcing Go to resolve dependencies
COPY . .  # Copy the full source now (since go.sum doesn't exist yet)

# Generate go.sum and download modules
RUN go mod tidy

# Build the app
RUN package controllers

import (
    "encoding/json"
    "net/http"
    "path/filepath"
    "strings"

    "github.com/gin-gonic/gin"
    "bus-booking/config"
    "bus-booking/models"
    "mime/multipart"
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

// POST /admin/buses
func CreateBus(c *gin.Context) {
    var input BusInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    amenitiesJSON, _ := json.Marshal(input.Amenities)

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

// Upload image for a bus
func UploadBusImage(c *gin.Context) {
    busID := c.Param("bus_id")

    file, err := c.FormFile("image")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
        return
    }

    if !isValidImage(file) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image format"})
        return
    }

    filename := filepath.Base(file.Filename)
    filePath := "uploads/" + filename
    if err := c.SaveUploadedFile(file, filePath); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save image"})
        return
    }

    image := models.BusImage{
        BusID:    parseUint(busID),
        ImageURL: filePath,
    }

    if err := config.DB.Create(&image).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully", "image": image})
}

// Helper: check if image has a valid extension
func isValidImage(file *multipart.FileHeader) bool {
    ext := strings.ToLower(filepath.Ext(file.Filename))
    return ext == ".jpg" || ext == ".jpeg" || ext == ".png"
}

// Helper: parse string to uint
func parseUint(s string) uint {
    var id uint
    fmt.Sscanf(s, "%d", &id)
    return id
}


EXPOSE 8080

CMD ["./main"]
