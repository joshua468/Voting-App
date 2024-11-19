package database

import (
	"log"

	"github.com/joshua468/voting-app/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() {
	// Get database configuration
	dsn := config.GetDBConfig()

	// Connect to the database using GORM
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database initialized successfully")
}
