package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	PropertyID uint     `json:"property_id"`
	Property   Property `json:"property" gorm:"foreignKey:PropertyID"`
	ReviewerID uint     `json:"reviewer_id"`
	Reviewer   User     `json:"reviewer" gorm:"foreignKey:ReviewerID"`
	Rating     int      `json:"rating" gorm:"check:rating >= 1 AND rating <= 5"`
	Comment    string   `json:"comment" gorm:"type:text"`
	ReviewType string   `json:"review_type"` // "property", "agent", "landlord"
	IsVerified bool     `json:"is_verified" gorm:"default:false"`
}
