package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joshua468/voting-app/internal/config"
	"github.com/joshua468/voting-app/internal/controllers"
	"github.com/joshua468/voting-app/internal/database"
	"github.com/joshua468/voting-app/internal/models" // Import your models package

	"github.com/joshua468/voting-app/internal/middlewares"
)

func main() {

	// Load environment variables
	config.LoadEnv()

	// Initialize the database connection
	database.InitDB()

	// Run AutoMigrate to ensure the database schema matches the models
	err := database.DB.AutoMigrate(&models.User{}, &models.Election{}, &models.Aspirant{}, &models.Vote{}) // Add all models here
	if err != nil {
		log.Fatal("Error running migrations:", err)
	}

	// Initialize Gin router
	r := gin.Default()

	// Routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Admin routes
	adminRoutes := r.Group("/admin", middlewares.AuthMiddleware("admin"))
	{
		adminRoutes.POST("/election", controllers.CreateElection)
		adminRoutes.PUT("/election/:id/winner", controllers.DeclareWinner)
	}

	// Public routes
	r.GET("/elections", controllers.GetElections)

	// Voter routes
	voterRoutes := r.Group("/voter", middlewares.AuthMiddleware("voter"))
	{
		voterRoutes.POST("/vote", controllers.Vote)
	}

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server running on port:", port)
	r.Run(":" + port)
}
