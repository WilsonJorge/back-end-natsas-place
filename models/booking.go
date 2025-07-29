package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	UserID             uint       `json:"user_id" gorm:"not null"`
	PropertyID         uint       `json:"property_id" gorm:"not null"`
	CheckInDate        time.Time  `json:"check_in_date" gorm:"not null"`
	CheckOutDate       time.Time  `json:"check_out_date" gorm:"not null"`
	Adults             int        `json:"adults" gorm:"not null;default:1"`
	Children           int        `json:"children" gorm:"default:0"`
	Infants            int        `json:"infants" gorm:"default:0"`
	Pets               int        `json:"pets" gorm:"default:0"`
	TotalNights        int        `json:"total_nights" gorm:"not null"`
	PricePerNight      float64    `json:"price_per_night" gorm:"not null"`
	SubtotalAmount     float64    `json:"subtotal_amount" gorm:"not null"`
	CleaningFee        float64    `json:"cleaning_fee" gorm:"default:0"`
	ServiceFee         float64    `json:"service_fee" gorm:"default:0"`
	TaxAmount          float64    `json:"tax_amount" gorm:"default:0"`
	TotalAmount        float64    `json:"total_amount" gorm:"not null"`
	Status             string     `json:"status" gorm:"not null;default:'PENDING'"`
	PaymentStatus      string     `json:"payment_status" gorm:"not null;default:'PENDING'"`
	PaymentMethod      string     `json:"payment_method"`
	SpecialRequests    string     `json:"special_requests" gorm:"type:text"`
	GuestMessage       string     `json:"guest_message" gorm:"type:text"`
	CancellationReason string     `json:"cancellation_reason" gorm:"type:text"`
	CheckInTime        *time.Time `json:"check_in_time"`
	CheckOutTime       *time.Time `json:"check_out_time"`
	ConfirmedAt        *time.Time `json:"confirmed_at"`
	CancelledAt        *time.Time `json:"cancelled_at"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`

	// Relationships
	User     User      `json:"user" gorm:"foreignKey:UserID"`
	Property Property  `json:"property" gorm:"foreignKey:PropertyID"`
	Review   *Review   `json:"review" gorm:"foreignKey:BookingID"`
	Payments []Payment `json:"payments" gorm:"foreignKey:BookingID"`
}

// Booking Status Constants
const (
	BookingStatusPending   = "PENDING"
	BookingStatusConfirmed = "CONFIRMED"
	BookingStatusCheckedIn = "CHECKED_IN"
	BookingStatusCompleted = "COMPLETED"
	BookingStatusCancelled = "CANCELLED"
	BookingStatusRefunded  = "REFUNDED"
)

// Payment Status Constants
const (
	PaymentStatusPending  = "PENDING"
	PaymentStatusPaid     = "PAID"
	PaymentStatusFailed   = "FAILED"
	PaymentStatusRefunded = "REFUNDED"
	PaymentStatusPartial  = "PARTIAL"
)

type BookingCreateInput struct {
	PropertyID      uint   `json:"property_id" validate:"required"`
	CheckInDate     string `json:"check_in_date" validate:"required"`
	CheckOutDate    string `json:"check_out_date" validate:"required"`
	Adults          int    `json:"adults" validate:"required,min=1,max=16"`
	Children        int    `json:"children" validate:"min=0,max=16"`
	Infants         int    `json:"infants" validate:"min=0,max=5"`
	Pets            int    `json:"pets" validate:"min=0,max=5"`
	SpecialRequests string `json:"special_requests"`
	GuestMessage    string `json:"guest_message"`
}

type BookingUpdateInput struct {
	Status             string `json:"status" validate:"oneof=PENDING CONFIRMED CHECKED_IN COMPLETED CANCELLED REFUNDED"`
	PaymentStatus      string `json:"payment_status" validate:"oneof=PENDING PAID FAILED REFUNDED PARTIAL"`
	SpecialRequests    string `json:"special_requests"`
	CancellationReason string `json:"cancellation_reason"`
}

type BookingResponse struct {
	ID              uint              `json:"id"`
	UserID          uint              `json:"user_id"`
	PropertyID      uint              `json:"property_id"`
	CheckInDate     time.Time         `json:"check_in_date"`
	CheckOutDate    time.Time         `json:"check_out_date"`
	Adults          int               `json:"adults"`
	Children        int               `json:"children"`
	Infants         int               `json:"infants"`
	Pets            int               `json:"pets"`
	TotalNights     int               `json:"total_nights"`
	PricePerNight   float64           `json:"price_per_night"`
	SubtotalAmount  float64           `json:"subtotal_amount"`
	CleaningFee     float64           `json:"cleaning_fee"`
	ServiceFee      float64           `json:"service_fee"`
	TaxAmount       float64           `json:"tax_amount"`
	TotalAmount     float64           `json:"total_amount"`
	Status          string            `json:"status"`
	PaymentStatus   string            `json:"payment_status"`
	PaymentMethod   string            `json:"payment_method"`
	SpecialRequests string            `json:"special_requests"`
	GuestMessage    string            `json:"guest_message"`
	CheckInTime     *time.Time        `json:"check_in_time"`
	CheckOutTime    *time.Time        `json:"check_out_time"`
	ConfirmedAt     *time.Time        `json:"confirmed_at"`
	CancelledAt     *time.Time        `json:"cancelled_at"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
	User            *UserResponse     `json:"user,omitempty"`
	Property        *PropertyResponse `json:"property,omitempty"`
	Review          *ReviewResponse   `json:"review,omitempty"`
}

type BookingFilterInput struct {
	Status        string `json:"status"`
	PaymentStatus string `json:"payment_status"`
	PropertyID    uint   `json:"property_id"`
	UserID        uint   `json:"user_id"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
	Page          int    `json:"page"`
	Limit         int    `json:"limit"`
}

type BookingCalendarResponse struct {
	Date        string  `json:"date"`
	IsAvailable bool    `json:"is_available"`
	Price       float64 `json:"price"`
	MinStay     int     `json:"min_stay"`
}

type BookingAvailabilityInput struct {
	PropertyID   uint   `json:"property_id" validate:"required"`
	CheckInDate  string `json:"check_in_date" validate:"required"`
	CheckOutDate string `json:"check_out_date" validate:"required"`
	Adults       int    `json:"adults" validate:"required,min=1"`
	Children     int    `json:"children" validate:"min=0"`
	Infants      int    `json:"infants" validate:"min=0"`
	Pets         int    `json:"pets" validate:"min=0"`
}
