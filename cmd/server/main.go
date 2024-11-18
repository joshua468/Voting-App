package main

import (
	"log"
	"os"

	"github.com/joshua468/voting-app/config"
	"github.com/joshua468/voting-app/controllers"
	"github.com/joshua468/voting-app/database"
	"github.com/joshua468/voting-app/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Initialize database
	database.InitDB()

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
