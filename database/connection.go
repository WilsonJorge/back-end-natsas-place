package database

import (
	"log"
	"natsas-place/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectBD() {

	database_url := config.GetConfig().Database_url

	db, err := gorm.Open(postgres.Open(database_url), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database")
	}

	DB = db
}
