package models

import "gorm.io/gorm"

type PropertyImage struct {
	gorm.Model
	PropertyID  uint   `json:"property_id"`
	ImageURL    string `json:"image_url" gorm:"not null"`
	Description string `json:"description"`
	IsMain      bool   `json:"is_main" gorm:"default:false"`
	Order       int    `json:"order" gorm:"default:0"`
}
