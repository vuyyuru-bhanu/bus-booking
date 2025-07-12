package models

import (
    
    "golang.org/x/crypto/bcrypt"
    "time"
)

type User struct {
    ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    Email     string    `gorm:"type:varchar(255);unique;not null" json:"email"`
    Password  string    `gorm:"type:varchar(255);not null" json:"password"`
    Name      string    `gorm:"type:varchar(255)" json:"name"`
    Role      string    `gorm:"type:enum('user','admin');default:'user';not null" json:"role"`
    CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
    DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}

func (u *User) BeforeSave() error {
    if u.Password != "" {
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
        if err != nil {
            return err
        }
        u.Password = string(hashedPassword)
    }
    return nil
}