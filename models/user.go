package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name                string     `json:"name" gorm:"not null" validate:"required,min=2,max=100"`
	Email               string     `json:"email" gorm:"unique;not null" validate:"required,email"`
	Password            string     `json:"-" gorm:"not null" validate:"required,min=8"`
	Phone               string     `json:"phone" validate:"required"`
	DateOfBirth         time.Time  `json:"date_of_birth" gorm:"type:date" validate:"required"`
	Role                string     `json:"role" gorm:"not null;default:'USER'" validate:"required,oneof=ADMIN USER"`
	IsVerified          bool       `json:"is_verified" gorm:"default:false"`
	IsActive            bool       `json:"is_active" gorm:"default:true"`
	ProfileImage        string     `json:"profile_image"`
	Gender              string     `json:"gender" validate:"oneof=MALE FEMALE OTHER"`
	Nationality         string     `json:"nationality"`
	Language            string     `json:"language" gorm:"default:'pt'"`
	EmailVerifiedAt     *time.Time `json:"email_verified_at"`
	LastLoginAt         *time.Time `json:"last_login_at"`
	ResetPasswordToken  string     `json:"-"`
	ResetPasswordExpiry *time.Time `json:"-"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`

	// Relationships
	Bookings      []Booking      `json:"bookings" gorm:"foreignKey:UserID"`
	Reviews       []Review       `json:"reviews" gorm:"foreignKey:ReviewerID"`
	Favorites     []Favorite     `json:"favorites" gorm:"foreignKey:UserID"`
	Notifications []Notification `json:"notifications" gorm:"foreignKey:UserID"`
	Payments      []Payment      `json:"payments" gorm:"foreignKey:UserID"`
}

type UserResponse struct {
	ID              uint       `json:"id"`
	Name            string     `json:"name"`
	Email           string     `json:"email"`
	Phone           string     `json:"phone"`
	Role            string     `json:"role"`
	IsVerified      bool       `json:"is_verified"`
	IsActive        bool       `json:"is_active"`
	ProfileImage    string     `json:"profile_image"`
	Gender          string     `json:"gender"`
	Nationality     string     `json:"nationality"`
	Language        string     `json:"language"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	LastLoginAt     *time.Time `json:"last_login_at"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

type UserRegistrationInput struct {
	Name        string `json:"name" validate:"required,min=2,max=100"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=8"`
	Phone       string `json:"phone" validate:"required"`
	DateOfBirth string `json:"date_of_birth" validate:"required"`
	Gender      string `json:"gender" validate:"oneof=MALE FEMALE OTHER"`
	Nationality string `json:"nationality"`
	Language    string `json:"language"`
}

type UserLoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserUpdateInput struct {
	Name         string `json:"name" validate:"min=2,max=100"`
	Phone        string `json:"phone"`
	Gender       string `json:"gender" validate:"oneof=MALE FEMALE OTHER"`
	Nationality  string `json:"nationality"`
	Language     string `json:"language"`
	ProfileImage string `json:"profile_image"`
}

type PasswordResetInput struct {
	Email string `json:"email" validate:"required,email"`
}

type PasswordUpdateInput struct {
	Token       string `json:"token" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=8"`
}

type PasswordChangeInput struct {
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required,min=8"`
}
