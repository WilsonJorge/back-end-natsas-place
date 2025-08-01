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

	// Only run migrations if explicitly enabled
	if config.GetConfig().RunMigrations {
		log.Println("🔄 Running database migrations...")
		createdMigration(db)
		log.Println("✅ Database migrations completed")
	} else {
		log.Println("⏭️ Skipping database migrations (RUN_MIGRATIONS=false)")
	}

	Instance = db
}
