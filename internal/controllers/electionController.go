package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshua468/voting-app/internal/database"
	"github.com/joshua468/voting-app/internal/models"
	"github.com/joshua468/voting-app/internal/utils"
)

// CreateElection creates a new election with validation.
func CreateElection(c *gin.Context) {
	var election models.Election
	if err := c.ShouldBindJSON(&election); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid request payload: "+err.Error())
		return
	}

	// Ensure title is provided
	if election.Title == "" {
		utils.RespondJSON(c, http.StatusBadRequest, "Election title is required")
		return
	}
	var existingElection models.Election
	if err := database.DB.Where("title = ?", election.Title).First(&existingElection).Error; err == nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Election title must be unique")
		return
	}

	// Ensure status is valid ("active" or "closed")
	if election.Status != "active" && election.Status != "closed" {
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid election status. It must be 'active' or 'closed'")
		return
	}

	// Ensure at least one aspirant is present
	if len(election.Aspirants) == 0 {
		utils.RespondJSON(c, http.StatusBadRequest, "At least one aspirant is required for the election")
		return
	}

	// Ensure all aspirants are valid (check if they exist in the database)
	for _, aspirant := range election.Aspirants {
		var existingAspirant models.Aspirant
		if err := database.DB.Where("id = ?", aspirant.ID).First(&existingAspirant).Error; err != nil {
			utils.RespondJSON(c, http.StatusBadRequest, "Invalid aspirant ID: "+err.Error())
			return
		}
	}

	// Ensure the winner is valid (if provided, should be one of the aspirants)
	if election.WinnerID != 0 {
		var winner models.Aspirant
		if err := database.DB.Where("id = ?", election.WinnerID).First(&winner).Error; err != nil {
			utils.RespondJSON(c, http.StatusBadRequest, "Invalid winner ID: "+err.Error())
			return
		}
	}

	// Create the election in the database
	if err := database.DB.Create(&election).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Error creating election: "+err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusCreated, gin.H{
		"message":  "Election created successfully",
		"election": election,
	})
}

// GetElections retrieves all elections and their associated aspirants.
func GetElections(c *gin.Context) {
	var elections []models.Election
	// Preload Aspirants to eager load them with elections
	if err := database.DB.Preload("Aspirants").Find(&elections).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Error retrieving elections: "+err.Error())
		return
	}

	// Return the list of elections
	utils.RespondJSON(c, http.StatusOK, elections)
}
