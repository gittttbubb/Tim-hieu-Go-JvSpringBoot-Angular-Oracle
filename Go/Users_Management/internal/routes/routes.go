package routes

import (
	"github.com/gofiber/fiber/v2"

	"go-rbac-system/internal/handler"
	"go-rbac-system/internal/middleware"
)

type RouteConfig struct {
	AuthHandler       *handler.AuthHandler
	UserHandler       *handler.UserHandler
	RoleHandler       *handler.RoleHandler
	PermissionHandler *handler.PermissionHandler
	AuditHandler      *handler.AuditHandler

	AuthMiddleware       middleware.AuthMiddleware
	PermissionMiddleware middleware.PermissionMiddleware
}
func RegisterAllRoutes(app *fiber.App, cfg RouteConfig) {

	api := app.Group("/api/v1")

	RegisterAuthRoutes(api, cfg)
	RegisterUserRoutes(api, cfg)
	RegisterRoleRoutes(api, cfg)
	RegisterPermissionRoutes(api, cfg)
	RegisterAuditRoutes(api, cfg)
}
func RegisterAuthRoutes(api fiber.Router, cfg RouteConfig) {

	auth := api.Group("/auth")

	auth.Post("/login", cfg.AuthHandler.Login)

	auth.Post("/forgot-password", cfg.AuthHandler.ForgotPassword)
	auth.Post("/reset-password", cfg.AuthHandler.ResetPassword)

	auth.Post("/change-password",
		cfg.AuthMiddleware.RequireAuth(),
		cfg.AuthHandler.ChangePassword,
	)

	// Chưa dùng đến
	auth.Post("/refresh",
		cfg.AuthMiddleware.RequireAuth(),
		cfg.AuthHandler.Refresh,
	)
}

func RegisterUserRoutes(api fiber.Router, cfg RouteConfig) {

	users := api.Group("/users",
		cfg.AuthMiddleware.RequireAuth(),
	)

	users.Get("/",
		cfg.PermissionMiddleware.Require("USER_VIEW"),
		cfg.UserHandler.List,
	)

	users.Get("/:id",
		cfg.PermissionMiddleware.Require("USER_VIEW"),
		cfg.UserHandler.GetByID,
	)

	users.Post("/",
		cfg.PermissionMiddleware.Require("USER_CREATE"),
		cfg.UserHandler.Create,
	)

	users.Put("/:id",
		cfg.PermissionMiddleware.Require("USER_UPDATE"),
		cfg.UserHandler.Update,
	)

	users.Delete("/:id",
		cfg.PermissionMiddleware.Require("USER_DELETE"),
		cfg.UserHandler.Delete,
	)

	// USER STATUS
	users.Post("/:id/lock",
		cfg.PermissionMiddleware.Require("USER_UPDATE"),
		cfg.UserHandler.Lock,
	)

	users.Post("/:id/unlock",
		cfg.PermissionMiddleware.Require("USER_UPDATE"),
		cfg.UserHandler.Unlock,
	)

	// USER PERMISSION OVERRIDES (gộp trong UserHandler)
	users.Get("/:id/overrides",
		cfg.PermissionMiddleware.Require("USER_PERMISSION_VIEW"),
		cfg.UserHandler.GetOverrides,
	)

	users.Post("/:id/overrides/assign",
		cfg.PermissionMiddleware.Require("USER_PERMISSION_ASSIGN"),
		cfg.UserHandler.AssignOverride,
	)
	users.Delete("/overrides/:id",
    cfg.PermissionMiddleware.Require("USER_PERMISSION_REMOVE"),
    cfg.UserHandler.RemoveOverride,
)
}

func RegisterRoleRoutes(api fiber.Router, cfg RouteConfig) {

	roles := api.Group("/roles",
		cfg.AuthMiddleware.RequireAuth(),
	)

	roles.Get("/",
		cfg.PermissionMiddleware.Require("ROLE_VIEW"),
		cfg.RoleHandler.List,
	)

	roles.Get("/:id",
		cfg.PermissionMiddleware.Require("ROLE_VIEW"),
		cfg.RoleHandler.GetByID,
	)

	roles.Post("/",
		cfg.PermissionMiddleware.Require("ROLE_CREATE"),
		cfg.RoleHandler.Create,
	)

	roles.Put("/:id",
		cfg.PermissionMiddleware.Require("ROLE_UPDATE"),
		cfg.RoleHandler.Update,
	)

	roles.Delete("/:id",
		cfg.PermissionMiddleware.Require("ROLE_DELETE"),
		cfg.RoleHandler.Delete,
	)

	// ROLE PERMISSIONS (gộp trong RoleHandler)
	roles.Get("/:id/permissions",
		cfg.PermissionMiddleware.Require("ROLE_PERMISSION_VIEW"),
		cfg.RoleHandler.GetPermissions,
	)

	roles.Post("/permissions/assign",
		cfg.PermissionMiddleware.Require("ROLE_PERMISSION_ASSIGN"),
		cfg.RoleHandler.AssignPermission,
	)

	roles.Delete("/:id/permissions/:permissionId",
		cfg.PermissionMiddleware.Require("ROLE_PERMISSION_REMOVE"),
		cfg.RoleHandler.RemovePermission,
	)
}

func RegisterPermissionRoutes(api fiber.Router, cfg RouteConfig) {

	permissions := api.Group("/permissions",
		cfg.AuthMiddleware.RequireAuth(),
	)

	permissions.Get("/",
		cfg.PermissionMiddleware.Require("PERMISSION_VIEW"),
		cfg.PermissionHandler.List,
	)

	permissions.Get("/:id",
		cfg.PermissionMiddleware.Require("PERMISSION_VIEW"),
		cfg.PermissionHandler.GetByID,
	)

	permissions.Get("/feature/:featureCode",
		cfg.PermissionMiddleware.Require("PERMISSION_VIEW"),
		cfg.PermissionHandler.GetByFeatureCode,
	)

	permissions.Get("/grouped",
		cfg.PermissionMiddleware.Require("PERMISSION_VIEW"),
		cfg.PermissionHandler.Grouped,
	)
}

func RegisterAuditRoutes(api fiber.Router, cfg RouteConfig) {

	audits := api.Group("/audits",
		cfg.AuthMiddleware.RequireAuth(),
	)

	audits.Get("/",
		cfg.PermissionMiddleware.Require("AUDIT_VIEW"),
		cfg.AuditHandler.ListWithFilter,
	)

	audits.Get("/:id",
		cfg.PermissionMiddleware.Require("AUDIT_VIEW"),
		cfg.AuditHandler.GetByID,
	)

	audits.Get("/actor/:actorId",
		cfg.PermissionMiddleware.Require("AUDIT_VIEW"),
		cfg.AuditHandler.ListByActor,
	)

	audits.Get("/entity",
		cfg.PermissionMiddleware.Require("AUDIT_VIEW"),
		cfg.AuditHandler.ListByEntity,
	)

	audits.Get("/tenant",
		cfg.PermissionMiddleware.Require("AUDIT_VIEW"),
		cfg.AuditHandler.ListByTenant,
	)
}

