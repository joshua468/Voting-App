package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshua468/voting-app/internal/database"
	"github.com/joshua468/voting-app/internal/models"
	"github.com/joshua468/voting-app/internal/utils"
)

func DeclareWinner(c *gin.Context) {
	var election models.Election
	electionID := c.Param("id")

	if err := database.DB.Preload("Aspirants").First(&election, electionID).Error; err != nil {
		utils.RespondJSON(c, http.StatusNotFound, "Election not found")
		return
	}

	var winner models.Aspirant
	maxVotes := -1
	for _, aspirant := range election.Aspirants {
		if aspirant.Votes > maxVotes {
			winner = aspirant
			maxVotes = aspirant.Votes
		}
	}

	utils.RespondJSON(c, http.StatusOK, gin.H{
		"election": election.Title,
		"winner":   winner.Name,
	})
}
