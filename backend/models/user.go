package models

import (
	"errors"
	"time"

	"bus-booking/config"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Email     string     `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password  string     `gorm:"type:varchar(255);not null" json:"-"`
	Name      string     `gorm:"type:varchar(255)" json:"name"`
	Role      string     `gorm:"type:enum('user','admin');default:'user';not null" json:"role"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}

// HashPassword hashes the user password
func (u *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

// CheckPassword verifies the given password with the hashed one
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

import (
	"errors"
	"time"

	"bus-booking/config"
	"gorm.io/gorm"
)

// GetUserByID returns a user by ID
func GetUserByID(db *gorm.DB, userID int) (*User, error) {
	var user User
	result := db.First(&user, userID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	return &user, result.Error
}

// UpdateUserProfile updates the user's name and email
func UpdateUserProfile(db *gorm.DB, userID int, name string, email string) error {
	var user User
	result := db.First(&user, userID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("user not found")
	}
	user.Name = name
	user.Email = email
	return db.Save(&user).Error
}
