package controllers

import (
	"net/http"

	"github.com/joshua468/voting-app/models"

	"github.com/gin-gonic/gin"
)

// CreateElection - handles creation of elections
func CreateElection(c *gin.Context) {
	var election models.Election
	if err := c.ShouldBindJSON(&election); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Create election
	if err := models.CreateElection(&election); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create election"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Election created successfully"})
}

// GetElections - retrieves all elections
func GetElections(c *gin.Context) {
	elections, err := models.GetAllElections()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve elections"})
		return
	}

	c.JSON(http.StatusOK, elections)
}

// GetElectionByID - retrieves a single election by ID
func GetElectionByID(c *gin.Context) {
	id := c.Param("id")

	election, err := models.GetElectionByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Election not found"})
		return
	}

	c.JSON(http.StatusOK, election)
}
