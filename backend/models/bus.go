package models

import (
	"github.com/jinzhu/gorm"
)

type Bus struct {
	gorm.Model
	RouteID       uint        `gorm:"not null"`
	Company       string      `gorm:"not null"`
	AC            bool        `gorm:"not null"`
	Type          string      `gorm:"type:ENUM('seater','sleeper');not null"`
	Capacity      int         `gorm:"not null"`
	AvailableSeats int        `gorm:"not null"`
	Price         float64     `gorm:"type:DECIMAL(10,2);not null"`
	Amenities     string      `gorm:"type:JSON"`
	Images        []BusImage  `gorm:"foreignKey:BusID"`
	Route         Route       `gorm:"foreignKey:RouteID"`
}