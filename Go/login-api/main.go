package main

import (
	"log"

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
	r := routes.Setup(authHandler,)
	r.Run(":" + cfg.Port,)
}
