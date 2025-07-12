package models

import (
    "time"
)

type Bus struct {
    ID             uint       `gorm:"primaryKey;autoIncrement" json:"id"`
    RouteID        uint       `gorm:"not null" json:"route_id"`
    Company        string     `gorm:"not null" json:"company"`
    AC             bool       `gorm:"not null" json:"ac"`
    Type           string     `gorm:"type:ENUM('seater','sleeper');not null" json:"type"`
    Capacity       int        `gorm:"not null" json:"capacity"`
    AvailableSeats int        `gorm:"not null" json:"available_seats"`
    Price          float64    `gorm:"type:DECIMAL(10,2);not null" json:"price"`
    Amenities      string     `gorm:"type:JSON" json:"amenities"`
    CreatedAt      time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt      time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
    DeletedAt      *time.Time `gorm:"index" json:"deleted_at"`

    // Relationships
    Images []BusImage `gorm:"foreignKey:BusID" json:"images,omitempty"`
    Route  Route      `gorm:"foreignKey:RouteID" json:"route,omitempty"`
}
