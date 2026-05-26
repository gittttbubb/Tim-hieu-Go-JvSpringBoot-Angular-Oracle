package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"login-api/internal/utils"
)

func Auth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "missing token",
			})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		userId, err := utils.ParseToken(secret, tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
			})
			c.Abort()
			return
		}

		// gắn vào context
		c.Set("userId", userId)

		c.Next()
	}
}
