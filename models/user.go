package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string    `json:"name" gorm:"not null"`
	Email        string    `json:"email" gorm:"unique;not null"`
	Password     string    `json:"-" gorm:"not null"`
	Phone        string    `json:"phone"`
	DateOfBirth  time.Time `json:"date_of_birth"`
	Role         string    `json:"role" gorm:"not null"` // ADMIN, USER
	IsVerified   bool      `json:"is_verified" gorm:"default:false"`
	ProfileImage string    `json:"profile_image"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	Reviews   []Review   `json:"reviews" gorm:"foreignKey:ReviewerID"`
	Visits    []Visit    `json:"visits" gorm:"foreignKey:VisitorID"`
	Favorites []Favorite `json:"favorites" gorm:"foreignKey:UserID"`
}
