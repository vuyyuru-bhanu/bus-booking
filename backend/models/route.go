package models

import (
	"github.com/jinzhu/gorm"
)

type Route struct {
	gorm.Model
	Origin        string `gorm:"not null"`
	Destination   string `gorm:"not null"`
	DepartureTime string `gorm:"not null"`
	Duration      int    `gorm:"not null"`
}