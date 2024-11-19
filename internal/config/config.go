package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads the environment variables from the .env file.
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

// GetJWTSecret retrieves the JWT secret from environment variables.
func GetJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET not set in .env")
	}
	return secret
}

// GetDBConfig retrieves the database connection string from environment variables.
func GetDBConfig() string {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Check if any database-related environment variable is missing.
	if dbUser == "" || dbPassword == "" || dbHost == "" || dbPort == "" || dbName == "" {
		log.Fatal("One or more database environment variables are missing")
	}

	// Format the connection string to pass to the database driver.
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
}
