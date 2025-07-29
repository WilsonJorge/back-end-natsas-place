package database

import (
	"natsas-place/models"

	"gorm.io/gorm"
)

func createdMigration(dbInstance *gorm.DB) {
	// User model
	err := dbInstance.AutoMigrate(&models.User{})
	if err != nil {
		return
	}

	// Address model
	err = dbInstance.AutoMigrate(&models.Address{})
	if err != nil {
		return
	}

	// Property model
	err = dbInstance.AutoMigrate(&models.Property{})
	if err != nil {
		return
	}

	// PropertyImage model
	err = dbInstance.AutoMigrate(&models.PropertyImage{})
	if err != nil {
		return
	}

	// Amenity model
	err = dbInstance.AutoMigrate(&models.Amenity{})
	if err != nil {
		return
	}

	// PropertyAmenity model
	err = dbInstance.AutoMigrate(&models.PropertyAmenity{})
	if err != nil {
		return
	}

	// PropertyAvailability model
	err = dbInstance.AutoMigrate(&models.PropertyAvailability{})
	if err != nil {
		return
	}

	// PropertyPricing model
	err = dbInstance.AutoMigrate(&models.PropertyPricing{})
	if err != nil {
		return
	}

	// Booking model
	err = dbInstance.AutoMigrate(&models.Booking{})
	if err != nil {
		return
	}

	// Payment model
	err = dbInstance.AutoMigrate(&models.Payment{})
	if err != nil {
		return
	}

	// Review model
	err = dbInstance.AutoMigrate(&models.Review{})
	if err != nil {
		return
	}

	// Favorite model
	err = dbInstance.AutoMigrate(&models.Favorite{})
	if err != nil {
		return
	}

	// Notification model
	err = dbInstance.AutoMigrate(&models.Notification{})
	if err != nil {
		return
	}

	// NotificationPreferences model
	err = dbInstance.AutoMigrate(&models.NotificationPreferences{})
	if err != nil {
		return
	}

	// Create unique indexes
	createUniqueIndexes(dbInstance)
}

func createUniqueIndexes(dbInstance *gorm.DB) {
	// PropertyAmenity unique index
	dbInstance.Exec("CREATE UNIQUE INDEX IF NOT EXISTS idx_property_amenity_unique ON property_amenities(property_id, amenity_id)")

	// PropertyAvailability unique index
	dbInstance.Exec("CREATE UNIQUE INDEX IF NOT EXISTS idx_property_availability_unique ON property_availabilities(property_id, date)")

	// User email unique index (already handled by gorm tag)

	// NotificationPreferences unique index (already handled by gorm tag)
}
