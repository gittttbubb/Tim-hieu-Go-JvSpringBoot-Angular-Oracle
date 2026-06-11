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
		tokenInfo, err := utils.ParseToken(secret, tokenString)
		// Kiểm tra có lỗi không
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
			})
			c.Abort()
			return
		}

		c.Set("userId", tokenInfo.UserID)
		c.Set("role", tokenInfo.Role)

		c.Next()
	}
}
func Role(role string) gin.HandlerFunc {
    return func(c *gin.Context) {

        currentRole := c.GetString("role")

        if currentRole != role {
            c.JSON(http.StatusForbidden, gin.H{
                "message": "forbidden",
            })
            c.Abort()
            return
        }

        c.Next()
    }
}