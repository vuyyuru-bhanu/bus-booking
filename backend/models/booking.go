package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	ID         uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     uint       `gorm:"not null" json:"user_id"`
	BusID      uint       `gorm:"not null" json:"bus_id"`
	SeatNumber int        `gorm:"not null" json:"seat_number"`
	Status     string     `gorm:"type:enum('confirmed','cancelled');default:'confirmed'" json:"status"`
	CreatedAt  time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt  *time.Time `gorm:"index" json:"deleted_at"`

	// Associations for Preload
	User *User `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Bus  *Bus  `gorm:"foreignKey:BusID" json:"bus,omitempty"`
}

// GetBookingsByUserID returns all bookings for a given user ID
func GetBookingsByUserID(userID int) ([]Booking, error) {
	var bookings []Booking
	result := db.Preload("Bus").Preload("User").Where("user_id = ?", userID).Find(&bookings)
	return bookings, result.Error
}

// GetBookingByIDAndUserID returns a booking by ID and user ID
func GetBookingByIDAndUserID(bookingID string, userID int) (*Booking, error) {
	var booking Booking
	result := db.Preload("Bus").Preload("User").Where("id = ? AND user_id = ?", bookingID, userID).First(&booking)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("booking not found")
	}
	return &booking, result.Error
}
