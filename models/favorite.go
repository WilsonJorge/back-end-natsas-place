package models

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	UserID     uint     `json:"user_id"`
	User       User     `json:"user" gorm:"foreignKey:UserID"`
	PropertyID uint     `json:"property_id"`
	Property   Property `json:"property" gorm:"foreignKey:PropertyID"`
}
