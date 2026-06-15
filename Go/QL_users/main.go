package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"

	"go-user-management/internal/config"
	"go-user-management/internal/database"
	"go-user-management/internal/handler"
	"go-user-management/internal/middleware"
	"go-user-management/internal/repository"
	"go-user-management/internal/routes"
	"go-user-management/internal/service"
)

func main() {
	// Load Configuration
	cfg, err := config.Load()
	log.Printf(
		"Environment: %s",
		os.Getenv("APP_ENV"),
	)
	if err != nil {
		log.Fatalf(
			"failed to load config: %v",
			err,
		)
	}

	// Connect Oracle Database
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf(
			"failed to connect database: %v",
			err,
		)
	}
	defer db.Close()

	// Repositories
	authRepo := repository.NewAuthRepository(db)
	userRepo := repository.NewUserRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	permissionRepo := repository.NewPermissionRepository(db)
	rolePermissionRepo := repository.NewRolePermissionRepository(db)
	overrideRepo := repository.NewUserPermissionOverrideRepository(db)
	passwordResetRepo := repository.NewPasswordResetRepository(db)
	auditRepo := repository.NewAuditRepository(db)

	// Services
	permissionResolverService := service.NewPermissionResolverService(userRepo, permissionRepo, overrideRepo,)
	authService := service.NewAuthService(authRepo, permissionResolverService, cfg,)
	userService := service.NewUserService(userRepo, roleRepo,)
	roleService := service.NewRoleService(roleRepo,)
	permissionService := service.NewPermissionService(permissionRepo)
	rolePermissionService := service.NewRolePermissionService(rolePermissionRepo, roleRepo, permissionRepo,)
	userOverrideService := service.NewUserPermissionOverrideService(overrideRepo, userRepo, permissionRepo,)
	passwordResetService := service.NewPasswordResetService(authRepo, userRepo, passwordResetRepo,)
	auditService :=service.NewAuditService(auditRepo,)

	// Handlers
	authHandler := handler.NewAuthHandler(authService,)
	userHandler := handler.NewUserHandler(userService,)
	roleHandler := handler.NewRoleHandler(roleService,)
	permissionHandler := handler.NewPermissionHandler(permissionService,)
	rolePermissionHandler := handler.NewRolePermissionHandler(rolePermissionService,)
	userOverrideHandler := handler.NewUserPermissionOverrideHandler(userOverrideService,)
	passwordResetHandler := handler.NewPasswordResetHandler(passwordResetService,)
	auditHandler := handler.NewAuditHandler(auditService,)

	// Middlewares

	authMiddleware := middleware.NewAuthMiddleware(cfg.JWT.Secret,)
	permissionMiddleware := middleware.NewPermissionMiddleware()

	// Fiber App
	app := fiber.New()

	// Routes
	routes.RegisterRoutes(
		app,
		routes.RouteConfig{
			AuthHandler:           authHandler,
			UserHandler:           userHandler,
			RoleHandler:           roleHandler,
			PermissionHandler:     permissionHandler,
			RolePermissionHandler: rolePermissionHandler,
			UserOverrideHandler:   userOverrideHandler,
			PasswordResetHandler:  passwordResetHandler,
			AuditHandler:          auditHandler,

			AuthMiddleware:       authMiddleware,
			PermissionMiddleware: permissionMiddleware,
		},
	)

	// Start Server
	log.Printf(
		"Server started on port %s",
		cfg.Server.Port,
	)
	if err := app.Listen(
		":" + cfg.Server.Port,
	); err != nil {
		log.Fatal(err)
	}
}