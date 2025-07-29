package models

import (
	"time"

	"gorm.io/gorm"
)

type PropertyAvailability struct {
	gorm.Model
	PropertyID    uint      `json:"property_id" gorm:"not null"`
	Date          time.Time `json:"date" gorm:"not null;type:date"`
	IsAvailable   bool      `json:"is_available" gorm:"default:true"`
	Price         float64   `json:"price" gorm:"not null"`
	MinStayNights int       `json:"min_stay_nights" gorm:"default:1"`
	MaxStayNights int       `json:"max_stay_nights" gorm:"default:30"`
	Notes         string    `json:"notes"`

	// Relationships
	Property Property `json:"property" gorm:"foreignKey:PropertyID"`
}

type PropertyPricing struct {
	gorm.Model
	PropertyID    uint      `json:"property_id" gorm:"not null"`
	StartDate     time.Time `json:"start_date" gorm:"not null;type:date"`
	EndDate       time.Time `json:"end_date" gorm:"not null;type:date"`
	Price         float64   `json:"price" gorm:"not null"`
	PriceType     string    `json:"price_type" gorm:"not null;default:'SEASONAL'"`
	Description   string    `json:"description"`
	MinStayNights int       `json:"min_stay_nights"`
	MaxStayNights int       `json:"max_stay_nights"`
	IsActive      bool      `json:"is_active" gorm:"default:true"`

	// Relationships
	Property Property `json:"property" gorm:"foreignKey:PropertyID"`
}

// Price Type Constants
const (
	PriceTypeBase     = "BASE"
	PriceTypeSeasonal = "SEASONAL"
	PriceTypeWeekend  = "WEEKEND"
	PriceTypeHoliday  = "HOLIDAY"
	PriceTypeSpecial  = "SPECIAL"
)

type PropertyAvailabilityCreateInput struct {
	PropertyID    uint    `json:"property_id" validate:"required"`
	Date          string  `json:"date" validate:"required"`
	IsAvailable   bool    `json:"is_available"`
	Price         float64 `json:"price" validate:"required,gt=0"`
	MinStayNights int     `json:"min_stay_nights" validate:"min=1,max=90"`
	MaxStayNights int     `json:"max_stay_nights" validate:"min=1,max=365"`
	Notes         string  `json:"notes"`
}

type PropertyAvailabilityUpdateInput struct {
	IsAvailable   bool    `json:"is_available"`
	Price         float64 `json:"price" validate:"gt=0"`
	MinStayNights int     `json:"min_stay_nights" validate:"min=1,max=90"`
	MaxStayNights int     `json:"max_stay_nights" validate:"min=1,max=365"`
	Notes         string  `json:"notes"`
}

type PropertyAvailabilityResponse struct {
	ID            uint    `json:"id"`
	PropertyID    uint    `json:"property_id"`
	Date          string  `json:"date"`
	IsAvailable   bool    `json:"is_available"`
	Price         float64 `json:"price"`
	MinStayNights int     `json:"min_stay_nights"`
	MaxStayNights int     `json:"max_stay_nights"`
	Notes         string  `json:"notes"`
}

type PropertyPricingCreateInput struct {
	PropertyID    uint    `json:"property_id" validate:"required"`
	StartDate     string  `json:"start_date" validate:"required"`
	EndDate       string  `json:"end_date" validate:"required"`
	Price         float64 `json:"price" validate:"required,gt=0"`
	PriceType     string  `json:"price_type" validate:"required,oneof=BASE SEASONAL WEEKEND HOLIDAY SPECIAL"`
	Description   string  `json:"description"`
	MinStayNights int     `json:"min_stay_nights" validate:"min=1,max=90"`
	MaxStayNights int     `json:"max_stay_nights" validate:"min=1,max=365"`
	IsActive      bool    `json:"is_active"`
}

type PropertyPricingUpdateInput struct {
	StartDate     string  `json:"start_date"`
	EndDate       string  `json:"end_date"`
	Price         float64 `json:"price" validate:"gt=0"`
	PriceType     string  `json:"price_type" validate:"oneof=BASE SEASONAL WEEKEND HOLIDAY SPECIAL"`
	Description   string  `json:"description"`
	MinStayNights int     `json:"min_stay_nights" validate:"min=1,max=90"`
	MaxStayNights int     `json:"max_stay_nights" validate:"min=1,max=365"`
	IsActive      bool    `json:"is_active"`
}

type PropertyPricingResponse struct {
	ID            uint    `json:"id"`
	PropertyID    uint    `json:"property_id"`
	StartDate     string  `json:"start_date"`
	EndDate       string  `json:"end_date"`
	Price         float64 `json:"price"`
	PriceType     string  `json:"price_type"`
	Description   string  `json:"description"`
	MinStayNights int     `json:"min_stay_nights"`
	MaxStayNights int     `json:"max_stay_nights"`
	IsActive      bool    `json:"is_active"`
}

type AvailabilityCalendarInput struct {
	PropertyID uint   `json:"property_id" validate:"required"`
	StartDate  string `json:"start_date" validate:"required"`
	EndDate    string `json:"end_date" validate:"required"`
}

type AvailabilityCalendarResponse struct {
	Date        string  `json:"date"`
	IsAvailable bool    `json:"is_available"`
	Price       float64 `json:"price"`
	MinStay     int     `json:"min_stay"`
	MaxStay     int     `json:"max_stay"`
	IsBooked    bool    `json:"is_booked"`
}

type BulkAvailabilityUpdateInput struct {
	PropertyID    uint                              `json:"property_id" validate:"required"`
	StartDate     string                            `json:"start_date" validate:"required"`
	EndDate       string                            `json:"end_date" validate:"required"`
	IsAvailable   *bool                             `json:"is_available"`
	Price         *float64                          `json:"price" validate:"omitempty,gt=0"`
	MinStayNights *int                              `json:"min_stay_nights" validate:"omitempty,min=1,max=90"`
	MaxStayNights *int                              `json:"max_stay_nights" validate:"omitempty,min=1,max=365"`
	Notes         *string                           `json:"notes"`
	ExcludeDates  []string                          `json:"exclude_dates"`
	SpecificDates []PropertyAvailabilityCreateInput `json:"specific_dates"`
}
