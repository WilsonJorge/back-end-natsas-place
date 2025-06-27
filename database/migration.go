package database

import (
	"natsas-place/models"

	"gorm.io/gorm"
)

func createtMigration(dbInstance *gorm.DB) {
	dbInstance.AutoMigrate(&models.User{})
	dbInstance.AutoMigrate(&models.Property{})
	dbInstance.AutoMigrate(&models.PropertyImage{})
	dbInstance.AutoMigrate(&models.Favorite{})
	dbInstance.AutoMigrate(&models.Visit{})
	dbInstance.AutoMigrate(&models.Review{})
	dbInstance.AutoMigrate(&models.Address{})
}
