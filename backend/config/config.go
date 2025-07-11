package config

import (
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"bus-booking/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic("Failed to connect to database")
	}

	database.AutoMigrate(&models.User{}, &models.Bus{}, &models.Route{}, &models.Booking{}, &models.BusImage{}, &models.Notification{})
	DB = database
}