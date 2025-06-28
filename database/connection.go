package database

import (
	"log"
	"natsas-place/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB

func ConnectBD() {

	databaseUrl := config.GetConfig().DatabaseUrl

	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database")
	}

	createdMigration(db)

	Instance = db
}
