package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

// AmenitiesArray é um tipo customizado para lidar com amenities como JSON
type AmenitiesArray []string

// Value implementa driver.Valuer para converter para JSON
func (a AmenitiesArray) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}
	return json.Marshal(a)
}

// Scan implementa sql.Scanner para converter de JSON
func (a *AmenitiesArray) Scan(value interface{}) error {
	if value == nil {
		*a = nil
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, a)
}

// Location represents the property location
type Location struct {
	Address     string `json:"address" gorm:"not null" validate:"required"`
	City        string `json:"city" gorm:"not null" validate:"required"`
	State       string `json:"state" gorm:"not null" validate:"required"`
	ZipCode     string `json:"zip_code" gorm:"not null" validate:"required"`
	Coordinates struct {
		Lat float64 `json:"lat" gorm:"not null" validate:"required"`
		Lng float64 `json:"lng" gorm:"not null" validate:"required"`
	} `json:"coordinates" gorm:"embedded"`
}

// PropertyImage represents an image of the property
type PropertyImage struct {
	gorm.Model
	PropertyID uint   `json:"property_id" gorm:"not null"`
	URL        string `json:"url" gorm:"not null" validate:"required,url"`
	Alt        string `json:"alt"`
	IsPrimary  bool   `json:"is_primary" gorm:"default:false"`
}

// BookedPeriod represents a period when the property is booked
type BookedPeriod struct {
	gorm.Model
	PropertyID uint      `json:"property_id" gorm:"not null"`
	StartDate  time.Time `json:"start_date" gorm:"not null" validate:"required"`
	EndDate    time.Time `json:"end_date" gorm:"not null" validate:"required"`
	Guest      User      `json:"guest" gorm:"foreignKey:GuestID"`
	GuestID    uint      `json:"guest_id" gorm:"not null"`
	Status     string    `json:"status" gorm:"not null;default:'pending'" validate:"required,oneof=pending confirmed cancelled completed"`
}

// Property represents a property listing
type Property struct {
	gorm.Model
	ID          uint            `json:"id"`
	Title       string          `json:"title" gorm:"not null" validate:"required,min=2,max=200"`
	Description string          `json:"description" gorm:"type:text" validate:"required,min=10"`
	Price       float64         `json:"price" gorm:"not null" validate:"required,min=0"`
	Location    Location        `json:"location" gorm:"embedded"`
	PriceType   string          `json:"price_type" gorm:"not null;default:'daily'" validate:"required,oneof=daily weekly monthly"`
	Bedrooms    int             `json:"bedrooms" gorm:"not null" validate:"required,min=0"`
	Bathrooms   int             `json:"bathrooms" gorm:"not null" validate:"required,min=0"`
	Area        float64         `json:"area" gorm:"not null" validate:"required,min=0"`
	Type        string          `json:"type" gorm:"not null" validate:"required,oneof=house apartment studio condo"`
	Amenities   AmenitiesArray  `json:"amenities" gorm:"type:json"`
	Images      []PropertyImage `json:"images" gorm:"foreignKey:PropertyID"`
	OwnerID     uint            `json:"owner_id" gorm:"not null"`
	Owner       User            `json:"owner" gorm:"foreignKey:OwnerID"`
	Status      string          `json:"status" gorm:"not null;default:'available'" validate:"required,oneof=available rented maintenance"`
	Available   bool            `json:"available" gorm:"default:true"`
	BookedDates []BookedPeriod  `json:"booked_dates" gorm:"foreignKey:PropertyID"`
	Featured    bool            `json:"featured" gorm:"default:false"`
	Reviews     []Review        `json:"reviews" gorm:"foreignKey:PropertyID"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

// PropertyCreateInput represents the input for creating a new property
type PropertyCreateInput struct {
	Title       string         `json:"title" validate:"required,min=2,max=200"`
	Description string         `json:"description" validate:"required,min=10"`
	Price       float64        `json:"price" validate:"required,min=0"`
	Location    Location       `json:"location" validate:"required"`
	PriceType   string         `json:"price_type" validate:"required,oneof=daily weekly monthly"`
	Bedrooms    int            `json:"bedrooms" validate:"required,min=0"`
	Bathrooms   int            `json:"bathrooms" validate:"required,min=0"`
	Area        float64        `json:"area" validate:"required,min=0"`
	Type        string         `json:"type" validate:"required,oneof=house apartment studio condo"`
	Amenities   AmenitiesArray `json:"amenities"`
	Images      []string       `json:"images" validate:"omitempty,dive,url"`
}

// PropertyUpdateInput represents the input for updating a property
type PropertyUpdateInput struct {
	Title       string         `json:"title" validate:"omitempty,min=2,max=200"`
	Description string         `json:"description" validate:"omitempty,min=10"`
	Price       *float64       `json:"price" validate:"omitempty,min=0"`
	Location    *Location      `json:"location"`
	PriceType   string         `json:"price_type" validate:"omitempty,oneof=daily weekly monthly"`
	Bedrooms    *int           `json:"bedrooms" validate:"omitempty,min=0"`
	Bathrooms   *int           `json:"bathrooms" validate:"omitempty,min=0"`
	Area        *float64       `json:"area" validate:"omitempty,min=0"`
	Type        string         `json:"type" validate:"omitempty,oneof=house apartment studio condo"`
	Amenities   AmenitiesArray `json:"amenities"`
	Status      string         `json:"status" validate:"omitempty,oneof=available rented maintenance"`
	Available   *bool          `json:"available"`
	Featured    *bool          `json:"featured"`
}

// PropertySearchInput represents the input for searching properties
type PropertySearchInput struct {
	Location     string         `json:"location"`
	MinPrice     *float64       `json:"min_price"`
	MaxPrice     *float64       `json:"max_price"`
	PriceType    string         `json:"price_type" validate:"omitempty,oneof=daily weekly monthly"`
	MinBedrooms  *int           `json:"min_bedrooms"`
	MaxBedrooms  *int           `json:"max_bedrooms"`
	MinBathrooms *int           `json:"min_bathrooms"`
	MaxBathrooms *int           `json:"max_bathrooms"`
	MinArea      *float64       `json:"min_area"`
	MaxArea      *float64       `json:"max_area"`
	Type         string         `json:"type" validate:"omitempty,oneof=house apartment studio condo"`
	Amenities    AmenitiesArray `json:"amenities"`
	Available    *bool          `json:"available"`
	Featured     *bool          `json:"featured"`
	StartDate    *time.Time     `json:"start_date"`
	EndDate      *time.Time     `json:"end_date"`
	Limit        int            `json:"limit" validate:"omitempty,min=1,max=100"`
	Offset       int            `json:"offset" validate:"omitempty,min=0"`
	OwnerID      *uint          `json:"owner_id"`
	Latitude     *float64       `json:"latitude"`
	Longitude    *float64       `json:"longitude"`
}

// PropertyImageInput represents the input for adding property images
type PropertyImageInput struct {
	URL       string `json:"url" validate:"required,url"`
	Alt       string `json:"alt"`
	IsPrimary bool   `json:"is_primary"`
}

// BookedPeriodInput represents the input for booking a property
type BookedPeriodInput struct {
	StartDate time.Time `json:"start_date" validate:"required"`
	EndDate   time.Time `json:"end_date" validate:"required"`
	GuestID   uint      `json:"guest_id" validate:"required"`
}

type Review struct {
	gorm.Model
	PropertyID uint   `json:"property_id" gorm:"not null"`
	Rating     int    `json:"rating" gorm:"not null" validate:"required,min=1,max=5"`
	Comment    string `json:"comment" gorm:"type:text" validate:"required,min=10"`
	ReviewerID uint   `json:"reviewer_id" gorm:"not null"`
	Reviewer   User   `json:"reviewer" gorm:"foreignKey:ReviewerID"`
}
