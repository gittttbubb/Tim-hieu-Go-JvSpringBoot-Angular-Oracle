package routes

import (
	"github.com/gin-gonic/gin"

	"login-api/internal/config"
	"login-api/internal/handler"
	"login-api/internal/middleware"
	"login-api/internal/service"
)

func Setup(auth *handler.AuthHandler, authService service.AuthService, cfg config.Config, r *gin.Engine) *gin.Engine {

	api := r.Group("/api")

	api.POST("/auth/login", auth.Login)
	
	api.POST("/auth/register", auth.Register)

	user := api.Group("/users")
	user.Use(middleware.Auth(cfg.JWTSecret))
	user.GET("/me", func(c *gin.Context) {
		// 1. lấy userId từ middleware
		userID := c.GetInt("userId")
		// 2. gọi service
		user, err := authService.GetUserByID(userID)
		if err != nil {
			c.JSON(404, gin.H{
				"message": "user not found",
			})
			return
		}
		
		// 3. trả response
		c.JSON(200, gin.H{
			"id":       user.ID,
			"username": user.Username,
			"role":     user.Role,

		})
		// Tìm cách trả về theo kiểu object, không phải map[string]interface{}
	})
	user.PUT("/profile", auth.UpdateProfile)
	user.PUT("/change-password", auth.ChangePassword)

	admin := api.Group("/admin")
	admin.Use(
		middleware.Auth(cfg.JWTSecret),
		middleware.Role("ADMIN"),
	)
	admin.GET("/list-users", auth.GetAllUsers)
	admin.PUT("/users/promote/:id", auth.PromoteToAdmin)
	
	return r
}