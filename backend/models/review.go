package models

import (
	"errors"
	"time"

	"bus-booking/config"
	"gorm.io/gorm"
)

type Review struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint       `gorm:"not null" json:"user_id"`
	BusID     uint       `gorm:"not null" json:"bus_id"`
	Rating    int        `gorm:"not null" json:"rating"`
	Comment   string     `gorm:"type:text" json:"comment"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`

	User *User `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Bus  *Bus  `gorm:"foreignKey:BusID" json:"bus,omitempty"`
}

// GetReviewsByBusID returns all reviews for a given bus ID
func GetReviewsByBusID(busID int) ([]Review, error) {
	var reviews []Review
	result := config.DB.Preload("User").Where("bus_id = ?", busID).Find(&reviews)
	return reviews, result.Error
}

// AddReview adds a new review
func AddReview(review *Review) error {
	result := config.DB.Create(review)
	return result.Error
}
