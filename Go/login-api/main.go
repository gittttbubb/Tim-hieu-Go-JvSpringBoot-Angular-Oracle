package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"login-api/internal/config"
	"login-api/internal/database"
	"login-api/internal/handler"
	"login-api/internal/repository"
	"login-api/internal/routes"
	"login-api/internal/service"
)

func main() {

	godotenv.Load()

	cfg := config.Load()

	db, err := database.Connect(cfg,)
	if err != nil {
		log.Fatal(err)
	}
	repo :=	repository.NewAuthRepository(db,)
	authService := service.NewAuthService(repo, cfg.JWTSecret,)
	authHandler := handler.NewAuthHandler(authService,)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:           12 * time.Hour,
	}))

	r = routes.Setup(authHandler,authService, cfg, r)
	r.Run(":" + cfg.Port,)
}
