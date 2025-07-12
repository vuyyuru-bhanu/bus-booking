package models

import (
	"time"
)

type Notification struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Message   string     `gorm:"type:text;not null" json:"message"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}
