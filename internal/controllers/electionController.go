package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshua468/voting-app/internal/database"
	"github.com/joshua468/voting-app/internal/models"
	"github.com/joshua468/voting-app/internal/utils"
)

func CreateElection(c *gin.Context) {
	var election models.Election
	if err := c.ShouldBindJSON(&election); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := database.DB.Create(&election).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Error creating election")
		return
	}

	utils.RespondJSON(c, http.StatusCreated, "Election created successfully")
}

func GetElections(c *gin.Context) {
	var elections []models.Election
	if err := database.DB.Preload("Aspirants").Find(&elections).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Error retrieving elections")
		return
	}

	c.JSON(http.StatusOK, elections)
}
