package database

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	migratePostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joshua468/voting-app/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection and applies migrations
func InitDB() {
	// Get database configuration
	dsn := config.GetDBConfig()

	// Connect to the database using GORM
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Retrieve *sql.DB from GORM
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get sql.DB from GORM:", err)
	}

	// Run migrations using the migrate tool
	RunMigrations(sqlDB, dsn)

	log.Println("Database initialized and migrations applied successfully")
}

// RunMigrations applies migrations from the migrations folder
func RunMigrations(sqlDB *sql.DB, dsn string) {
	// Create a migration driver from the *sql.DB
	driver, err := migratePostgres.WithInstance(sqlDB, &migratePostgres.Config{})
	if err != nil {
		log.Fatal("Failed to create migration driver:", err)
	}

	// Create a new migration instance
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal("Failed to create migration instance:", err)
	}

	// Apply migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migrations applied successfully")
}
