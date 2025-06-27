package models

import "gorm.io/gorm"

type Property struct {
	gorm.Model
	Title           string  `json:"title" gorm:"not null"`
	Description     string  `json:"description" gorm:"type:text"`
	PropertyType    string  `json:"property_type"`    // "house", "apartment", "land", "commercial"
	TransactionType string  `json:"transaction_type"` // "sale", "rent", "both"
	Price           float64 `json:"price" gorm:"not null"`
	RentPrice       float64 `json:"rent_price"`
	Area            float64 `json:"area"` // m²
	Bedrooms        int     `json:"bedrooms"`
	Bathrooms       int     `json:"bathrooms"`
	ParkingSpaces   int     `json:"parking_spaces"`
	Floor           int     `json:"floor"`
	TotalFloors     int     `json:"total_floors"`
	YearBuilt       int     `json:"year_built"`
	Status          string  `json:"status"`                    // "available", "rented", "sold", "inactive"
	Features        string  `json:"features" gorm:"type:text"` // JSON string

	AddressID uint            `json:"address_id"`
	Address   Address         `json:"address" gorm:"foreignKey:AddressID"`
	Images    []PropertyImage `json:"images" gorm:"foreignKey:PropertyID"`
	Favorites []Favorite      `json:"favorites" gorm:"foreignKey:PropertyID"`
}
