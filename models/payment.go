package models

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	BookingID       uint       `json:"booking_id" gorm:"not null"`
	UserID          uint       `json:"user_id" gorm:"not null"`
	Amount          float64    `json:"amount" gorm:"not null"`
	Currency        string     `json:"currency" gorm:"not null;default:'USD'"`
	PaymentMethod   string     `json:"payment_method" gorm:"not null"`
	PaymentProvider string     `json:"payment_provider" gorm:"not null"`
	TransactionID   string     `json:"transaction_id" gorm:"unique"`
	ExternalID      string     `json:"external_id"`
	Status          string     `json:"status" gorm:"not null;default:'PENDING'"`
	PaymentType     string     `json:"payment_type" gorm:"not null"`
	Description     string     `json:"description"`
	ProcessedAt     *time.Time `json:"processed_at"`
	FailureReason   string     `json:"failure_reason"`
	RefundAmount    float64    `json:"refund_amount" gorm:"default:0"`
	RefundReason    string     `json:"refund_reason"`
	RefundedAt      *time.Time `json:"refunded_at"`
	Metadata        string     `json:"metadata" gorm:"type:text"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`

	// Relationships
	Booking Booking `json:"booking" gorm:"foreignKey:BookingID"`
	User    User    `json:"user" gorm:"foreignKey:UserID"`
}

// Payment Status Constants
const (
	PaymentStatusProcessing = "PROCESSING"
	PaymentStatusCompleted  = "COMPLETED"
	PaymentStatusCancelled  = "CANCELLED"
)

// Payment Type Constants
const (
	PaymentTypeBooking = "BOOKING"
	PaymentTypeRefund  = "REFUND"
	PaymentTypeFee     = "FEE"
)

// Payment Method Constants
const (
	PaymentMethodCreditCard   = "CREDIT_CARD"
	PaymentMethodDebitCard    = "DEBIT_CARD"
	PaymentMethodPayPal       = "PAYPAL"
	PaymentMethodBankTransfer = "BANK_TRANSFER"
	PaymentMethodPIX          = "PIX"
)

// Payment Provider Constants
const (
	PaymentProviderStripe      = "STRIPE"
	PaymentProviderPayPal      = "PAYPAL"
	PaymentProviderMercadoPago = "MERCADO_PAGO"
)

type PaymentCreateInput struct {
	BookingID       uint    `json:"booking_id" validate:"required"`
	Amount          float64 `json:"amount" validate:"required,gt=0"`
	Currency        string  `json:"currency" validate:"required,len=3"`
	PaymentMethod   string  `json:"payment_method" validate:"required"`
	PaymentProvider string  `json:"payment_provider" validate:"required"`
	Description     string  `json:"description"`
}

type PaymentUpdateInput struct {
	Status        string `json:"status" validate:"oneof=PENDING PROCESSING COMPLETED FAILED CANCELLED REFUNDED"`
	TransactionID string `json:"transaction_id"`
	ExternalID    string `json:"external_id"`
	FailureReason string `json:"failure_reason"`
}

type PaymentResponse struct {
	ID              uint       `json:"id"`
	BookingID       uint       `json:"booking_id"`
	UserID          uint       `json:"user_id"`
	Amount          float64    `json:"amount"`
	Currency        string     `json:"currency"`
	PaymentMethod   string     `json:"payment_method"`
	PaymentProvider string     `json:"payment_provider"`
	TransactionID   string     `json:"transaction_id"`
	Status          string     `json:"status"`
	PaymentType     string     `json:"payment_type"`
	Description     string     `json:"description"`
	ProcessedAt     *time.Time `json:"processed_at"`
	FailureReason   string     `json:"failure_reason"`
	RefundAmount    float64    `json:"refund_amount"`
	RefundReason    string     `json:"refund_reason"`
	RefundedAt      *time.Time `json:"refunded_at"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

type RefundInput struct {
	PaymentID   uint    `json:"payment_id" validate:"required"`
	Amount      float64 `json:"amount" validate:"required,gt=0"`
	Reason      string  `json:"reason" validate:"required"`
	Description string  `json:"description"`
}

type PaymentFilterInput struct {
	Status          string  `json:"status"`
	PaymentMethod   string  `json:"payment_method"`
	PaymentProvider string  `json:"payment_provider"`
	BookingID       uint    `json:"booking_id"`
	UserID          uint    `json:"user_id"`
	StartDate       string  `json:"start_date"`
	EndDate         string  `json:"end_date"`
	MinAmount       float64 `json:"min_amount"`
	MaxAmount       float64 `json:"max_amount"`
	Page            int     `json:"page"`
	Limit           int     `json:"limit"`
}
