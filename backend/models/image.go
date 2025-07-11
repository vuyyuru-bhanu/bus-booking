package models

import (
	"github.com/jinzhu/gorm"
)

type BusImage struct {
	gorm.Model
	BusID    uint   `gorm:"not null"`
	ImageURL string `gorm:"not null"`
}