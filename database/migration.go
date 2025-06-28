package database

import (
	"natsas-place/models"

	"gorm.io/gorm"
)

func createdMigration(dbInstance *gorm.DB) {
	err := dbInstance.AutoMigrate(&models.User{})
	if err != nil {
		return
	}

	err = dbInstance.AutoMigrate(&models.Property{})
	if err != nil {
		return
	}

	err = dbInstance.AutoMigrate(&models.PropertyImage{})
	if err != nil {
		return
	}
	
	err = dbInstance.AutoMigrate(&models.Favorite{})
	if err != nil {
		return
	}

	err = dbInstance.AutoMigrate(&models.Visit{})
	if err != nil {
		return
	}

	err = dbInstance.AutoMigrate(&models.Review{})
	if err != nil {
		return
	}

	err = dbInstance.AutoMigrate(&models.Address{})
	if err != nil {
		return
	}
}
