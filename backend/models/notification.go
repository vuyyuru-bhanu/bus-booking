package models

import (
	"github.com/jinzhu/gorm"
)

type Notification struct {
	gorm.Model
	Message string `gorm:"type:TEXT;not null"`
}