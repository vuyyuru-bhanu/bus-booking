package models

import "time"

type BusImage struct {
    ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
    BusID     uint       `gorm:"not null" json:"bus_id"`
    ImageURL  string     `gorm:"type:varchar(255);not null" json:"image_url"`
    CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
    DeletedAt *time.Time `gorm:"index" json:"deleted_at"`

    Bus Bus `gorm:"foreignKey:BusID" json:"-"`
}
