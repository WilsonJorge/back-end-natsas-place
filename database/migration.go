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
