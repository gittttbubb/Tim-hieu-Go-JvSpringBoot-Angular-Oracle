package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"login-api/internal/utils"
)
// Check ROlE trong token 
func Auth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		// Kiểm tra có bắt đầu chuỗi bằng "Bearer " hay không
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "missing token",
			})
			// Dừng hệ thống, không cho request đi tiếp
			c.Abort()
			return	
		}
		// strings.TrimPrefix loại bỏ phần "Bearer " giữ lại phần token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		userId, err := utils.ParseToken(secret, tokenString)
		// Kiểm tra có lỗi không
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
