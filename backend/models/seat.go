package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Seat struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	BusID     uint       `gorm:"not null" json:"bus_id"`
	SeatNumber int       `gorm:"not null" json:"seat_number"`
	IsBooked  bool       `gorm:"default:false" json:"is_booked"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}

// GetAvailableSeats returns all available seats for a given bus ID
func GetAvailableSeats(busID int) ([]Seat, error) {
	var seats []Seat
	result := db.Where("bus_id = ? AND is_booked = false", busID).Find(&seats)
	return seats, result.Error
}

// SelectSeats marks seats as booked for a given bus ID and seat numbers
func SelectSeats(busID int, seatNumbers []int) error {
	tx := db.Begin()
	for _, seatNumber := range seatNumbers {
		result := tx.Model(&Seat{}).Where("bus_id = ? AND seat_number = ? AND is_booked = false", busID, seatNumber).Update("is_booked", true)
		if result.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("seat not available")
		}
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	}
	return tx.Commit().Error
}
