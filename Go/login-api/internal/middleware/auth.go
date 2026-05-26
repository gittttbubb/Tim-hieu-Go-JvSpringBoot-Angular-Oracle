package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("Authorization")
		// Kiểm tra Authorization bắt đầu bằng chữ "Bearer " hay không. Chưa thực sự giải mã (Decode/Verify) xem Token đó có hợp lệ hay đã hết hạn
		if !strings.HasPrefix(token, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})

			c.Abort()
			return
		}

		c.Next()
	}
}