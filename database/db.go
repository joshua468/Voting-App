package database

import (
    "fmt"
    "log"
    "os"
    "github.com/joshua468/voting-app/models"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Migrate the models to the database
    err = DB.AutoMigrate(&models.Election{}, &models.Aspirant{}, &models.Vote{})
    if err != nil {
        log.Fatal("Error migrating models:", err)
    }
}
