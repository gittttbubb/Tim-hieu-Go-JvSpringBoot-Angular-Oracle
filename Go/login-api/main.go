package main

import (
	// "os"

	"github.com/joho/godotenv"

	"login-api/internal/config"
	"login-api/internal/handler"
	"login-api/internal/repository"
	"login-api/internal/routes"
	"login-api/internal/service"
)

func main() {
	// Đọc file .env
	godotenv.Load() 

	cfg := config.Load()

	repo := repository.NewAuthRepository()

	authService := service.NewAuthService(
		repo,
		cfg.JWTSecret,
	)

	authHandler := handler.NewAuthHandler(
		authService,
	)

	r := routes.Setup(
		authHandler,
	)

	r.Run(":" + cfg.Port)

}
