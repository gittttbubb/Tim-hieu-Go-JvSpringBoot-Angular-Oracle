package bootstrap

import (
	"database/sql"

	"go-rbac-system/internal/config"
	"go-rbac-system/internal/handler"
	"go-rbac-system/internal/middleware"
	"go-rbac-system/internal/repository"
	"go-rbac-system/internal/service"
)

type Container struct {
	Config *config.Config
	DB     *sql.DB
	AuthMiddleware       middleware.AuthMiddleware
	PermissionMiddleware middleware.PermissionMiddleware
	AuthHandler       *handler.AuthHandler
	UserHandler       *handler.UserHandler
	RoleHandler       *handler.RoleHandler
	PermissionHandler *handler.PermissionHandler
	AuditHandler      *handler.AuditHandler
}

func NewContainer(cfg *config.Config, db *sql.DB) *Container {
	// repositories
	authRepo := repository.NewAuthRepository(db)
	userRepo := repository.NewUserRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	permissionRepo := repository.NewPermissionRepository(db)
	rolePermissionRepo := repository.NewRolePermissionRepository(db)
	userOverrideRepo := repository.NewUserPermissionOverrideRepository(db)
	auditRepo := repository.NewAuditRepository(db)
	passwordResetRepo := repository.NewPasswordResetRepository(db)


	// services
	userService := service.NewUserService(userRepo, roleRepo, auditRepo)
	roleService := service.NewRoleService(roleRepo)
	permissionService := service.NewPermissionService(permissionRepo)
	rolePermissionService := service.NewRolePermissionService(
		rolePermissionRepo,
		roleRepo,
		permissionRepo,
		userOverrideRepo,
	)
	authService := service.NewAuthService(authRepo, roleRepo, auditRepo, rolePermissionService)
	userOverrideService := service.NewUserPermissionOverrideService(
		userOverrideRepo,
		userRepo,
		permissionRepo,
	)
	emailService := service.NewEmailService(
		cfg.SMTP.Host,
		cfg.SMTP.Port,
		cfg.SMTP.Username,
		cfg.SMTP.Password,
		cfg.SMTP.From,
	)
// 	err := emailService.Send(
// 	"nvthang7891011@gmail.com",
// 	"SMTP Test",
// 	"<h1>Hello</h1>",
// )

// 	if err != nil {
// 		log.Fatal(err)
// 	}
	passwordResetService := service.NewPasswordResetService(userRepo, passwordResetRepo, auditRepo, emailService, cfg.Frontend.URL,)
	auditService := service.NewAuditService(auditRepo)

	// middleware
	authMiddleware := middleware.NewAuthMiddleware(cfg)
	permissionMiddleware := middleware.NewPermissionMiddleware(authService, permissionRepo)

	// handlers
	return &Container{
		Config: cfg,
		DB:     db,
		AuthMiddleware:       authMiddleware,
		PermissionMiddleware: permissionMiddleware,
		AuthHandler:       handler.NewAuthHandler(authService, passwordResetService, cfg),
		UserHandler:       handler.NewUserHandler(userService, userOverrideService),
		RoleHandler:       handler.NewRoleHandler(roleService, rolePermissionService),
		PermissionHandler: handler.NewPermissionHandler(permissionService),
		AuditHandler:      handler.NewAuditHandler(auditService),
	}
}