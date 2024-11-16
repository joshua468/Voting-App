package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres" 
	"gorm.io/gorm"            
)

var DB *gorm.DB
func LoadEnv() {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// GetDBConfig returns the database connection string
func GetDBConfig() string {
	// Read database configuration from environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Return the PostgreSQL connection string
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)
}


func InitializeDatabase() {
	dsn := GetDBConfig()

	// Connect to the PostgreSQL database using GORM
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Log successful connection
	log.Println("Database connected successfully")
}

// GetJWTSecretKey returns the JWT secret key for signing and parsing tokens
func GetJWTSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET is required in the .env file")
	}
	return secret
}
