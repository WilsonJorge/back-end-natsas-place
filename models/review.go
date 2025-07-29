package models

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	PropertyID          uint      `json:"property_id" gorm:"not null"`
	BookingID           uint      `json:"booking_id" gorm:"not null"`
	ReviewerID          uint      `json:"reviewer_id" gorm:"not null"`
	OverallRating       int       `json:"overall_rating" gorm:"not null" validate:"required,min=1,max=5"`
	CleanlinessRating   int       `json:"cleanliness_rating" gorm:"not null" validate:"required,min=1,max=5"`
	AccuracyRating      int       `json:"accuracy_rating" gorm:"not null" validate:"required,min=1,max=5"`
	CheckInRating       int       `json:"check_in_rating" gorm:"not null" validate:"required,min=1,max=5"`
	CommunicationRating int       `json:"communication_rating" gorm:"not null" validate:"required,min=1,max=5"`
	LocationRating      int       `json:"location_rating" gorm:"not null" validate:"required,min=1,max=5"`
	ValueRating         int       `json:"value_rating" gorm:"not null" validate:"required,min=1,max=5"`
	Comment             string    `json:"comment" gorm:"type:text" validate:"required,min=10,max=2000"`
	PublicComment       string    `json:"public_comment" gorm:"type:text"`
	PrivateComment      string    `json:"private_comment" gorm:"type:text"`
	IsVisible           bool      `json:"is_visible" gorm:"default:true"`
	IsRecommended       bool      `json:"is_recommended" gorm:"default:true"`
	IsVerified          bool      `json:"is_verified" gorm:"default:false"`
	WouldStayAgain      bool      `json:"would_stay_again" gorm:"default:true"`
	ReviewedAt          time.Time `json:"reviewed_at" gorm:"autoCreateTime"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`

	// Relationships
	Property Property `json:"property" gorm:"foreignKey:PropertyID"`
	Booking  Booking  `json:"booking" gorm:"foreignKey:BookingID"`
	Reviewer User     `json:"reviewer" gorm:"foreignKey:ReviewerID"`
}

type ReviewCreateInput struct {
	PropertyID          uint   `json:"property_id" validate:"required"`
	BookingID           uint   `json:"booking_id" validate:"required"`
	OverallRating       int    `json:"overall_rating" validate:"required,min=1,max=5"`
	CleanlinessRating   int    `json:"cleanliness_rating" validate:"required,min=1,max=5"`
	AccuracyRating      int    `json:"accuracy_rating" validate:"required,min=1,max=5"`
	CheckInRating       int    `json:"check_in_rating" validate:"required,min=1,max=5"`
	CommunicationRating int    `json:"communication_rating" validate:"required,min=1,max=5"`
	LocationRating      int    `json:"location_rating" validate:"required,min=1,max=5"`
	ValueRating         int    `json:"value_rating" validate:"required,min=1,max=5"`
	Comment             string `json:"comment" validate:"required,min=10,max=2000"`
	PublicComment       string `json:"public_comment" validate:"max=2000"`
	PrivateComment      string `json:"private_comment" validate:"max=2000"`
	IsRecommended       bool   `json:"is_recommended"`
	WouldStayAgain      bool   `json:"would_stay_again"`
}

type ReviewUpdateInput struct {
	OverallRating       int    `json:"overall_rating" validate:"min=1,max=5"`
	CleanlinessRating   int    `json:"cleanliness_rating" validate:"min=1,max=5"`
	AccuracyRating      int    `json:"accuracy_rating" validate:"min=1,max=5"`
	CheckInRating       int    `json:"check_in_rating" validate:"min=1,max=5"`
	CommunicationRating int    `json:"communication_rating" validate:"min=1,max=5"`
	LocationRating      int    `json:"location_rating" validate:"min=1,max=5"`
	ValueRating         int    `json:"value_rating" validate:"min=1,max=5"`
	Comment             string `json:"comment" validate:"min=10,max=2000"`
	PublicComment       string `json:"public_comment" validate:"max=2000"`
	PrivateComment      string `json:"private_comment" validate:"max=2000"`
	IsVisible           bool   `json:"is_visible"`
	IsRecommended       bool   `json:"is_recommended"`
	WouldStayAgain      bool   `json:"would_stay_again"`
}

type ReviewResponse struct {
	ID                  uint              `json:"id"`
	PropertyID          uint              `json:"property_id"`
	BookingID           uint              `json:"booking_id"`
	ReviewerID          uint              `json:"reviewer_id"`
	OverallRating       int               `json:"overall_rating"`
	CleanlinessRating   int               `json:"cleanliness_rating"`
	AccuracyRating      int               `json:"accuracy_rating"`
	CheckInRating       int               `json:"check_in_rating"`
	CommunicationRating int               `json:"communication_rating"`
	LocationRating      int               `json:"location_rating"`
	ValueRating         int               `json:"value_rating"`
	Comment             string            `json:"comment"`
	PublicComment       string            `json:"public_comment"`
	IsVisible           bool              `json:"is_visible"`
	IsRecommended       bool              `json:"is_recommended"`
	IsVerified          bool              `json:"is_verified"`
	WouldStayAgain      bool              `json:"would_stay_again"`
	ReviewedAt          time.Time         `json:"reviewed_at"`
	CreatedAt           time.Time         `json:"created_at"`
	UpdatedAt           time.Time         `json:"updated_at"`
	Reviewer            *UserResponse     `json:"reviewer,omitempty"`
	Property            *PropertyResponse `json:"property,omitempty"`
	Booking             *BookingResponse  `json:"booking,omitempty"`
}

type ReviewFilterInput struct {
	PropertyID     uint   `json:"property_id"`
	ReviewerID     uint   `json:"reviewer_id"`
	MinRating      int    `json:"min_rating" validate:"min=1,max=5"`
	MaxRating      int    `json:"max_rating" validate:"min=1,max=5"`
	IsVisible      *bool  `json:"is_visible"`
	IsRecommended  *bool  `json:"is_recommended"`
	IsVerified     *bool  `json:"is_verified"`
	WouldStayAgain *bool  `json:"would_stay_again"`
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
	SortBy         string `json:"sort_by"`
	SortOrder      string `json:"sort_order"`
	Page           int    `json:"page"`
	Limit          int    `json:"limit"`
}

type ReviewStatsResponse struct {
	TotalReviews         int64   `json:"total_reviews"`
	AverageRating        float64 `json:"average_rating"`
	CleanlinessAverage   float64 `json:"cleanliness_average"`
	AccuracyAverage      float64 `json:"accuracy_average"`
	CheckInAverage       float64 `json:"check_in_average"`
	CommunicationAverage float64 `json:"communication_average"`
	LocationAverage      float64 `json:"location_average"`
	ValueAverage         float64 `json:"value_average"`
	Rating5Count         int64   `json:"rating_5_count"`
	Rating4Count         int64   `json:"rating_4_count"`
	Rating3Count         int64   `json:"rating_3_count"`
	Rating2Count         int64   `json:"rating_2_count"`
	Rating1Count         int64   `json:"rating_1_count"`
	RecommendedCount     int64   `json:"recommended_count"`
	WouldStayAgainCount  int64   `json:"would_stay_again_count"`
}

type ReviewReminderInput struct {
	BookingID uint `json:"booking_id" validate:"required"`
}

type ReviewModerationInput struct {
	ReviewID       uint   `json:"review_id" validate:"required"`
	IsVisible      bool   `json:"is_visible"`
	IsVerified     bool   `json:"is_verified"`
	ModerationNote string `json:"moderation_note"`
}
