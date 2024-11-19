package utils

import (
	"github.com/gin-gonic/gin"
)

func RespondJSON(c *gin.Context, status int, data interface{}) {
	c.JSON(status, gin.H{"message": data})
}
