package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshua468/voting-app/internal/database"
	"github.com/joshua468/voting-app/internal/models"
	"github.com/joshua468/voting-app/internal/utils"
)

func Vote(c *gin.Context) {
	var vote models.Vote
	if err := c.ShouldBindJSON(&vote); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, err.Error())
		return
	}

	// Check if the election exists and is active
	var election models.Election
	if err := database.DB.Where("id = ? AND start_date <= NOW() AND end_date >= NOW()", vote.ElectionID).First(&election).Error; err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid or expired election")
		return
	}

	// Check if the aspirant exists in the election
	var aspirant models.Aspirant
	if err := database.DB.Where("id = ? AND election_id = ?", vote.AspirantID, vote.ElectionID).First(&aspirant).Error; err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid aspirant for this election")
		return
	}

	// Check if user already voted in the election
	var existingVote models.Vote
	if err := database.DB.Where("user_id = ? AND election_id = ?", vote.UserID, vote.ElectionID).First(&existingVote).Error; err == nil {
		utils.RespondJSON(c, http.StatusConflict, "User has already voted")
		return
	}

	// Create the new vote
	if err := database.DB.Create(&vote).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Error submitting vote")
		return
	}

	utils.RespondJSON(c, http.StatusCreated, "Vote submitted successfully")
}
