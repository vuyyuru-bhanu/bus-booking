package models

import (
	"github.com/jinzhu/gorm"
)

type Booking struct {
	gorm.Model
	UserID     uint   `gorm:"not null"`
	BusID      uint   `gorm:"not null"`
	SeatNumber int    `gorm:"not null"`
	Status     string `gorm:"type:ENUM('confirmed','cancelled');default:'confirmed'"`
}