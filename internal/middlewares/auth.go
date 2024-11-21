package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joshua468/voting-app/internal/utils"
)

func AuthMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			utils.RespondJSON(c, http.StatusUnauthorized, "Missing or invalid authorization token")
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			utils.RespondJSON(c, http.StatusForbidden, "Invalid token: "+err.Error())
			c.Abort()
			return
		}

		if claims.Role != role {
			utils.RespondJSON(c, http.StatusForbidden, "Unauthorized for this resource")
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)
		c.Next()
	}
}
