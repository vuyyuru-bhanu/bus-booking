package models

import (
    
    "time"
)

type Booking struct {
    ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    UserID    uint      `gorm:"not null" json:"user_id"`
    BusID     uint      `gorm:"not null" json:"bus_id"`
    SeatNumber int      `gorm:"not null" json:"seat_number"`
    Status    string    `gorm:"type:enum('confirmed','cancelled');default:'confirmed'" json:"status"`
    CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
    DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}