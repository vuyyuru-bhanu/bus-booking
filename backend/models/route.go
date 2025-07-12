package models

import (
	"time"
)

type Route struct {
	ID            uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Origin        string     `gorm:"type:varchar(255);not null" json:"origin"`
	Destination   string     `gorm:"type:varchar(255);not null" json:"destination"`
	DepartureTime string     `gorm:"type:varchar(255);not null" json:"departure_time"` // Consider time.Time if parsing datetime
	Duration      int        `gorm:"not null" json:"duration"` // Duration in minutes
	CreatedAt     time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt     *time.Time `gorm:"index" json:"deleted_at,omitempty"`

	Buses         []Bus      `gorm:"foreignKey:RouteID" json:"buses,omitempty"`
}
