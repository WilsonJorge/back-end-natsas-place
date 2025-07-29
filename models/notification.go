package models

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	UserID     uint       `json:"user_id" gorm:"not null"`
	Type       string     `json:"type" gorm:"not null"`
	Title      string     `json:"title" gorm:"not null"`
	Message    string     `json:"message" gorm:"type:text;not null"`
	Status     string     `json:"status" gorm:"not null;default:'UNREAD'"`
	Channel    string     `json:"channel" gorm:"not null"`
	Priority   string     `json:"priority" gorm:"not null;default:'MEDIUM'"`
	BookingID  *uint      `json:"booking_id"`
	PropertyID *uint      `json:"property_id"`
	PaymentID  *uint      `json:"payment_id"`
	Data       string     `json:"data" gorm:"type:text"`
	SentAt     *time.Time `json:"sent_at"`
	ReadAt     *time.Time `json:"read_at"`
	ExternalID string     `json:"external_id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`

	// Relationships
	User     User      `json:"user" gorm:"foreignKey:UserID"`
	Booking  *Booking  `json:"booking" gorm:"foreignKey:BookingID"`
	Property *Property `json:"property" gorm:"foreignKey:PropertyID"`
	Payment  *Payment  `json:"payment" gorm:"foreignKey:PaymentID"`
}

// Notification Status Constants
const (
	NotificationStatusUnread = "UNREAD"
	NotificationStatusRead   = "READ"
	NotificationStatusSent   = "SENT"
	NotificationStatusFailed = "FAILED"
)

// Notification Type Constants
const (
	NotificationTypeBookingConfirmed  = "BOOKING_CONFIRMED"
	NotificationTypeBookingCancelled  = "BOOKING_CANCELLED"
	NotificationTypeBookingReminder   = "BOOKING_REMINDER"
	NotificationTypeCheckInReminder   = "CHECK_IN_REMINDER"
	NotificationTypeCheckOutReminder  = "CHECK_OUT_REMINDER"
	NotificationTypePaymentSuccessful = "PAYMENT_SUCCESSFUL"
	NotificationTypePaymentFailed     = "PAYMENT_FAILED"
	NotificationTypeReviewRequest     = "REVIEW_REQUEST"
	NotificationTypeWelcome           = "WELCOME"
	NotificationTypeEmailVerification = "EMAIL_VERIFICATION"
	NotificationTypePasswordReset     = "PASSWORD_RESET"
	NotificationTypeNewMessage        = "NEW_MESSAGE"
	NotificationTypeSystemMaintenance = "SYSTEM_MAINTENANCE"
	NotificationTypePromotionalOffer  = "PROMOTIONAL_OFFER"
)

// Notification Channel Constants
const (
	NotificationChannelEmail = "EMAIL"
	NotificationChannelSMS   = "SMS"
	NotificationChannelPush  = "PUSH"
	NotificationChannelInApp = "IN_APP"
)

// Notification Priority Constants
const (
	NotificationPriorityLow    = "LOW"
	NotificationPriorityMedium = "MEDIUM"
	NotificationPriorityHigh   = "HIGH"
	NotificationPriorityUrgent = "URGENT"
)

type NotificationCreateInput struct {
	UserID     uint   `json:"user_id" validate:"required"`
	Type       string `json:"type" validate:"required"`
	Title      string `json:"title" validate:"required"`
	Message    string `json:"message" validate:"required"`
	Channel    string `json:"channel" validate:"required"`
	Priority   string `json:"priority" validate:"oneof=LOW MEDIUM HIGH URGENT"`
	BookingID  *uint  `json:"booking_id"`
	PropertyID *uint  `json:"property_id"`
	PaymentID  *uint  `json:"payment_id"`
	Data       string `json:"data"`
}

type NotificationUpdateInput struct {
	Status  string `json:"status" validate:"oneof=UNREAD READ SENT FAILED"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

type NotificationResponse struct {
	ID         uint       `json:"id"`
	UserID     uint       `json:"user_id"`
	Type       string     `json:"type"`
	Title      string     `json:"title"`
	Message    string     `json:"message"`
	Status     string     `json:"status"`
	Channel    string     `json:"channel"`
	Priority   string     `json:"priority"`
	BookingID  *uint      `json:"booking_id"`
	PropertyID *uint      `json:"property_id"`
	PaymentID  *uint      `json:"payment_id"`
	SentAt     *time.Time `json:"sent_at"`
	ReadAt     *time.Time `json:"read_at"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

type NotificationFilterInput struct {
	Status     string `json:"status"`
	Type       string `json:"type"`
	Channel    string `json:"channel"`
	Priority   string `json:"priority"`
	UserID     uint   `json:"user_id"`
	BookingID  uint   `json:"booking_id"`
	PropertyID uint   `json:"property_id"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
}

type NotificationPreferences struct {
	gorm.Model
	UserID                   uint `json:"user_id" gorm:"not null;uniqueIndex"`
	BookingConfirmationEmail bool `json:"booking_confirmation_email" gorm:"default:true"`
	BookingConfirmationSMS   bool `json:"booking_confirmation_sms" gorm:"default:false"`
	BookingReminderEmail     bool `json:"booking_reminder_email" gorm:"default:true"`
	BookingReminderSMS       bool `json:"booking_reminder_sms" gorm:"default:false"`
	PaymentNotificationEmail bool `json:"payment_notification_email" gorm:"default:true"`
	PaymentNotificationSMS   bool `json:"payment_notification_sms" gorm:"default:false"`
	ReviewRequestEmail       bool `json:"review_request_email" gorm:"default:true"`
	ReviewRequestSMS         bool `json:"review_request_sms" gorm:"default:false"`
	PromotionalOffersEmail   bool `json:"promotional_offers_email" gorm:"default:true"`
	PromotionalOffersSMS     bool `json:"promotional_offers_sms" gorm:"default:false"`
	SystemMaintenanceEmail   bool `json:"system_maintenance_email" gorm:"default:true"`
	SystemMaintenanceSMS     bool `json:"system_maintenance_sms" gorm:"default:false"`

	// Relationships
	User User `json:"user" gorm:"foreignKey:UserID"`
}

type NotificationPreferencesInput struct {
	BookingConfirmationEmail bool `json:"booking_confirmation_email"`
	BookingConfirmationSMS   bool `json:"booking_confirmation_sms"`
	BookingReminderEmail     bool `json:"booking_reminder_email"`
	BookingReminderSMS       bool `json:"booking_reminder_sms"`
	PaymentNotificationEmail bool `json:"payment_notification_email"`
	PaymentNotificationSMS   bool `json:"payment_notification_sms"`
	ReviewRequestEmail       bool `json:"review_request_email"`
	ReviewRequestSMS         bool `json:"review_request_sms"`
	PromotionalOffersEmail   bool `json:"promotional_offers_email"`
	PromotionalOffersSMS     bool `json:"promotional_offers_sms"`
	SystemMaintenanceEmail   bool `json:"system_maintenance_email"`
	SystemMaintenanceSMS     bool `json:"system_maintenance_sms"`
}
