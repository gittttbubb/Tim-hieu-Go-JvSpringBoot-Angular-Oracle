package routes

import (
	"github.com/gin-gonic/gin"

	"login-api/internal/handler"
	"login-api/internal/middleware"
)

func Setup(auth *handler.AuthHandler) *gin.Engine {

	r := gin.Default()

	api := r.Group("/api")

	api.POST("/auth/login", auth.Login)

	user := api.Group("/users")
	user.Use(middleware.Auth())

	user.GET("/me", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": "admin",
		})
	})

	return r
}