package models

import (
	"gorm.io/gorm"
)

type Amenity struct {
	gorm.Model
	Name        string `json:"name" gorm:"unique;not null" validate:"required,min=2,max=100"`
	Description string `json:"description" gorm:"type:text"`
	Icon        string `json:"icon"`
	Category    string `json:"category" gorm:"not null" validate:"required,oneof=BASIC COMFORT SAFETY ENTERTAINMENT ACCESSIBILITY"`
	IsActive    bool   `json:"is_active" gorm:"default:true"`

	// Relationships
	PropertyAmenities []PropertyAmenity `json:"property_amenities" gorm:"foreignKey:AmenityID"`
}

type PropertyAmenity struct {
	gorm.Model
	PropertyID uint `json:"property_id" gorm:"not null"`
	AmenityID  uint `json:"amenity_id" gorm:"not null"`

	// Relationships
	Property Property `json:"property" gorm:"foreignKey:PropertyID"`
	Amenity  Amenity  `json:"amenity" gorm:"foreignKey:AmenityID"`
}

// Amenity Category Constants
const (
	AmenityCategoryBasic         = "BASIC"
	AmenityCategoryComfort       = "COMFORT"
	AmenityCategorySafety        = "SAFETY"
	AmenityCategoryEntertainment = "ENTERTAINMENT"
	AmenityCategoryAccessibility = "ACCESSIBILITY"
)

type AmenityCreateInput struct {
	Name        string `json:"name" validate:"required,min=2,max=100"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Category    string `json:"category" validate:"required,oneof=BASIC COMFORT SAFETY ENTERTAINMENT ACCESSIBILITY"`
	IsActive    bool   `json:"is_active"`
}

type AmenityUpdateInput struct {
	Name        string `json:"name" validate:"min=2,max=100"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Category    string `json:"category" validate:"oneof=BASIC COMFORT SAFETY ENTERTAINMENT ACCESSIBILITY"`
	IsActive    bool   `json:"is_active"`
}

type AmenityResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Category    string `json:"category"`
	IsActive    bool   `json:"is_active"`
}

type PropertyAmenityResponse struct {
	ID      uint            `json:"id"`
	Amenity AmenityResponse `json:"amenity"`
}

type AmenityFilterInput struct {
	Category string `json:"category"`
	IsActive *bool  `json:"is_active"`
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
}
