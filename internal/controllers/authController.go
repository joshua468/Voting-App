package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshua468/voting-app/internal/database"
	"github.com/joshua468/voting-app/internal/models"
	"github.com/joshua468/voting-app/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println("Error binding JSON:", err)
		utils.RespondJSON(c, http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	if err := database.DB.Create(&user).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "User already exists")
		return
	}

	utils.RespondJSON(c, http.StatusCreated, "User registered successfully")
}

func Login(c *gin.Context) {
	var input models.User
	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		utils.RespondJSON(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		utils.RespondJSON(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Error generating token")
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
