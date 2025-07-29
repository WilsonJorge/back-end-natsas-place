package models

import (
	"time"

	"gorm.io/gorm"
)

type Property struct {
	gorm.Model
	Title              string     `json:"title" gorm:"not null" validate:"required,min=5,max=100"`
	Description        string     `json:"description" gorm:"type:text" validate:"required,min=50,max=2000"`
	PropertyType       string     `json:"property_type" gorm:"not null" validate:"required,oneof=HOUSE APARTMENT VILLA STUDIO LOFT CHALET CABIN COTTAGE"`
	MaxGuests          int        `json:"max_guests" gorm:"not null" validate:"required,min=1,max=16"`
	Bedrooms           int        `json:"bedrooms" gorm:"not null" validate:"required,min=1,max=10"`
	Bathrooms          int        `json:"bathrooms" gorm:"not null" validate:"required,min=1,max=10"`
	Beds               int        `json:"beds" gorm:"not null" validate:"required,min=1,max=20"`
	Size               float64    `json:"size" validate:"min=10"` // m²
	BasePrice          float64    `json:"base_price" gorm:"not null" validate:"required,gt=0"`
	CleaningFee        float64    `json:"cleaning_fee" gorm:"default:0" validate:"gte=0"`
	ServiceFee         float64    `json:"service_fee" gorm:"default:0" validate:"gte=0"`
	TaxRate            float64    `json:"tax_rate" gorm:"default:0" validate:"gte=0,lte=100"`
	MinStayNights      int        `json:"min_stay_nights" gorm:"default:1" validate:"min=1,max=90"`
	MaxStayNights      int        `json:"max_stay_nights" gorm:"default:30" validate:"min=1,max=365"`
	CheckInTime        string     `json:"check_in_time" gorm:"default:'15:00'" validate:"required"`
	CheckOutTime       string     `json:"check_out_time" gorm:"default:'11:00'" validate:"required"`
	Status             string     `json:"status" gorm:"not null;default:'ACTIVE'" validate:"required,oneof=ACTIVE INACTIVE MAINTENANCE"`
	IsInstantBookable  bool       `json:"is_instant_bookable" gorm:"default:true"`
	AllowPets          bool       `json:"allow_pets" gorm:"default:false"`
	AllowSmoking       bool       `json:"allow_smoking" gorm:"default:false"`
	AllowEvents        bool       `json:"allow_events" gorm:"default:false"`
	AllowChildren      bool       `json:"allow_children" gorm:"default:true"`
	HasWifi            bool       `json:"has_wifi" gorm:"default:true"`
	HasParking         bool       `json:"has_parking" gorm:"default:false"`
	HasKitchen         bool       `json:"has_kitchen" gorm:"default:false"`
	HasAirConditioning bool       `json:"has_air_conditioning" gorm:"default:false"`
	HasHeating         bool       `json:"has_heating" gorm:"default:false"`
	HasSwimmingPool    bool       `json:"has_swimming_pool" gorm:"default:false"`
	HasGym             bool       `json:"has_gym" gorm:"default:false"`
	HasBalcony         bool       `json:"has_balcony" gorm:"default:false"`
	HasTerrace         bool       `json:"has_terrace" gorm:"default:false"`
	HasGarden          bool       `json:"has_garden" gorm:"default:false"`
	HasFireplace       bool       `json:"has_fireplace" gorm:"default:false"`
	HasTv              bool       `json:"has_tv" gorm:"default:false"`
	HasWashingMachine  bool       `json:"has_washing_machine" gorm:"default:false"`
	HasDishwasher      bool       `json:"has_dishwasher" gorm:"default:false"`
	HasElevator        bool       `json:"has_elevator" gorm:"default:false"`
	HasSecurity        bool       `json:"has_security" gorm:"default:false"`
	HasAccessibility   bool       `json:"has_accessibility" gorm:"default:false"`
	Instructions       string     `json:"instructions" gorm:"type:text"`
	HouseRules         string     `json:"house_rules" gorm:"type:text"`
	CancellationPolicy string     `json:"cancellation_policy" gorm:"not null;default:'MODERATE'" validate:"required,oneof=FLEXIBLE MODERATE STRICT"`
	LastMaintenanceAt  *time.Time `json:"last_maintenance_at"`
	FeaturedUntil      *time.Time `json:"featured_until"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`

	// Foreign Keys
	AddressID uint `json:"address_id" gorm:"not null"`

	// Relationships
	Address           Address                `json:"address" gorm:"foreignKey:AddressID"`
	Images            []PropertyImage        `json:"images" gorm:"foreignKey:PropertyID"`
	Amenities         []PropertyAmenity      `json:"amenities" gorm:"foreignKey:PropertyID"`
	Bookings          []Booking              `json:"bookings" gorm:"foreignKey:PropertyID"`
	Reviews           []Review               `json:"reviews" gorm:"foreignKey:PropertyID"`
	Favorites         []Favorite             `json:"favorites" gorm:"foreignKey:PropertyID"`
	AvailabilityRules []PropertyAvailability `json:"availability_rules" gorm:"foreignKey:PropertyID"`
	PricingRules      []PropertyPricing      `json:"pricing_rules" gorm:"foreignKey:PropertyID"`
	Notifications     []Notification         `json:"notifications" gorm:"foreignKey:PropertyID"`
}

// Property Type Constants
const (
	PropertyTypeHouse     = "HOUSE"
	PropertyTypeApartment = "APARTMENT"
	PropertyTypeVilla     = "VILLA"
	PropertyTypeStudio    = "STUDIO"
	PropertyTypeLoft      = "LOFT"
	PropertyTypeChalet    = "CHALET"
	PropertyTypeCabin     = "CABIN"
	PropertyTypeCottage   = "COTTAGE"
)

// Property Status Constants
const (
	PropertyStatusActive      = "ACTIVE"
	PropertyStatusInactive    = "INACTIVE"
	PropertyStatusMaintenance = "MAINTENANCE"
)

// Cancellation Policy Constants
const (
	CancellationPolicyFlexible = "FLEXIBLE"
	CancellationPolicyModerate = "MODERATE"
	CancellationPolicyStrict   = "STRICT"
)

type PropertyCreateInput struct {
	Title              string  `json:"title" validate:"required,min=5,max=100"`
	Description        string  `json:"description" validate:"required,min=50,max=2000"`
	PropertyType       string  `json:"property_type" validate:"required,oneof=HOUSE APARTMENT VILLA STUDIO LOFT CHALET CABIN COTTAGE"`
	MaxGuests          int     `json:"max_guests" validate:"required,min=1,max=16"`
	Bedrooms           int     `json:"bedrooms" validate:"required,min=1,max=10"`
	Bathrooms          int     `json:"bathrooms" validate:"required,min=1,max=10"`
	Beds               int     `json:"beds" validate:"required,min=1,max=20"`
	Size               float64 `json:"size" validate:"min=10"`
	BasePrice          float64 `json:"base_price" validate:"required,gt=0"`
	CleaningFee        float64 `json:"cleaning_fee" validate:"gte=0"`
	ServiceFee         float64 `json:"service_fee" validate:"gte=0"`
	TaxRate            float64 `json:"tax_rate" validate:"gte=0,lte=100"`
	MinStayNights      int     `json:"min_stay_nights" validate:"min=1,max=90"`
	MaxStayNights      int     `json:"max_stay_nights" validate:"min=1,max=365"`
	CheckInTime        string  `json:"check_in_time" validate:"required"`
	CheckOutTime       string  `json:"check_out_time" validate:"required"`
	IsInstantBookable  bool    `json:"is_instant_bookable"`
	AllowPets          bool    `json:"allow_pets"`
	AllowSmoking       bool    `json:"allow_smoking"`
	AllowEvents        bool    `json:"allow_events"`
	AllowChildren      bool    `json:"allow_children"`
	Instructions       string  `json:"instructions"`
	HouseRules         string  `json:"house_rules"`
	CancellationPolicy string  `json:"cancellation_policy" validate:"required,oneof=FLEXIBLE MODERATE STRICT"`

	// Address
	AddressInput AddressCreateInput `json:"address" validate:"required"`

	// Amenities
	AmenityIDs []uint `json:"amenity_ids"`

	// Images
	ImageInputs []PropertyImageCreateInput `json:"images" validate:"required,min=1"`
}

type PropertyUpdateInput struct {
	Title              string  `json:"title" validate:"min=5,max=100"`
	Description        string  `json:"description" validate:"min=50,max=2000"`
	PropertyType       string  `json:"property_type" validate:"oneof=HOUSE APARTMENT VILLA STUDIO LOFT CHALET CABIN COTTAGE"`
	MaxGuests          int     `json:"max_guests" validate:"min=1,max=16"`
	Bedrooms           int     `json:"bedrooms" validate:"min=1,max=10"`
	Bathrooms          int     `json:"bathrooms" validate:"min=1,max=10"`
	Beds               int     `json:"beds" validate:"min=1,max=20"`
	Size               float64 `json:"size" validate:"min=10"`
	BasePrice          float64 `json:"base_price" validate:"gt=0"`
	CleaningFee        float64 `json:"cleaning_fee" validate:"gte=0"`
	ServiceFee         float64 `json:"service_fee" validate:"gte=0"`
	TaxRate            float64 `json:"tax_rate" validate:"gte=0,lte=100"`
	MinStayNights      int     `json:"min_stay_nights" validate:"min=1,max=90"`
	MaxStayNights      int     `json:"max_stay_nights" validate:"min=1,max=365"`
	CheckInTime        string  `json:"check_in_time"`
	CheckOutTime       string  `json:"check_out_time"`
	Status             string  `json:"status" validate:"oneof=ACTIVE INACTIVE MAINTENANCE"`
	IsInstantBookable  bool    `json:"is_instant_bookable"`
	AllowPets          bool    `json:"allow_pets"`
	AllowSmoking       bool    `json:"allow_smoking"`
	AllowEvents        bool    `json:"allow_events"`
	AllowChildren      bool    `json:"allow_children"`
	Instructions       string  `json:"instructions"`
	HouseRules         string  `json:"house_rules"`
	CancellationPolicy string  `json:"cancellation_policy" validate:"oneof=FLEXIBLE MODERATE STRICT"`
}

type PropertyResponse struct {
	ID                 uint      `json:"id"`
	Title              string    `json:"title"`
	Description        string    `json:"description"`
	PropertyType       string    `json:"property_type"`
	MaxGuests          int       `json:"max_guests"`
	Bedrooms           int       `json:"bedrooms"`
	Bathrooms          int       `json:"bathrooms"`
	Beds               int       `json:"beds"`
	Size               float64   `json:"size"`
	BasePrice          float64   `json:"base_price"`
	CleaningFee        float64   `json:"cleaning_fee"`
	ServiceFee         float64   `json:"service_fee"`
	TaxRate            float64   `json:"tax_rate"`
	MinStayNights      int       `json:"min_stay_nights"`
	MaxStayNights      int       `json:"max_stay_nights"`
	CheckInTime        string    `json:"check_in_time"`
	CheckOutTime       string    `json:"check_out_time"`
	Status             string    `json:"status"`
	IsInstantBookable  bool      `json:"is_instant_bookable"`
	AllowPets          bool      `json:"allow_pets"`
	AllowSmoking       bool      `json:"allow_smoking"`
	AllowEvents        bool      `json:"allow_events"`
	AllowChildren      bool      `json:"allow_children"`
	Instructions       string    `json:"instructions"`
	HouseRules         string    `json:"house_rules"`
	CancellationPolicy string    `json:"cancellation_policy"`
	AverageRating      float64   `json:"average_rating"`
	ReviewCount        int       `json:"review_count"`
	BookingCount       int       `json:"booking_count"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`

	// Relationships
	Address   *AddressResponse          `json:"address,omitempty"`
	Images    []PropertyImageResponse   `json:"images,omitempty"`
	Amenities []PropertyAmenityResponse `json:"amenities,omitempty"`
	Reviews   []ReviewResponse          `json:"reviews,omitempty"`
}

type PropertyFilterInput struct {
	PropertyType       string  `json:"property_type"`
	Status             string  `json:"status"`
	MinPrice           float64 `json:"min_price"`
	MaxPrice           float64 `json:"max_price"`
	MinGuests          int     `json:"min_guests"`
	MaxGuests          int     `json:"max_guests"`
	MinBedrooms        int     `json:"min_bedrooms"`
	MaxBedrooms        int     `json:"max_bedrooms"`
	MinBathrooms       int     `json:"min_bathrooms"`
	MaxBathrooms       int     `json:"max_bathrooms"`
	CheckInDate        string  `json:"check_in_date"`
	CheckOutDate       string  `json:"check_out_date"`
	City               string  `json:"city"`
	Neighborhood       string  `json:"neighborhood"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	RadiusKm           float64 `json:"radius_km"`
	IsInstantBookable  *bool   `json:"is_instant_bookable"`
	AllowPets          *bool   `json:"allow_pets"`
	AllowSmoking       *bool   `json:"allow_smoking"`
	AllowEvents        *bool   `json:"allow_events"`
	AllowChildren      *bool   `json:"allow_children"`
	HasWifi            *bool   `json:"has_wifi"`
	HasParking         *bool   `json:"has_parking"`
	HasKitchen         *bool   `json:"has_kitchen"`
	HasAirConditioning *bool   `json:"has_air_conditioning"`
	HasSwimmingPool    *bool   `json:"has_swimming_pool"`
	AmenityIDs         []uint  `json:"amenity_ids"`
	MinRating          float64 `json:"min_rating"`
	SortBy             string  `json:"sort_by"`
	SortOrder          string  `json:"sort_order"`
	Page               int     `json:"page"`
	Limit              int     `json:"limit"`
}

type PropertySearchResponse struct {
	Properties []PropertyResponse `json:"properties"`
	Total      int64              `json:"total"`
	Page       int                `json:"page"`
	Limit      int                `json:"limit"`
	TotalPages int                `json:"total_pages"`
}
