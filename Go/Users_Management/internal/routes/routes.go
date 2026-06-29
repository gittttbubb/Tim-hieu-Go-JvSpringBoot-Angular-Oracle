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

	AuthRoutes(api, cfg)
	UserRoutes(api, cfg)
	RoleRoutes(api, cfg)
	PermissionRoutes(api, cfg)
	AuditRoutes(api, cfg)
}

// AUTH
func AuthRoutes(api fiber.Router, cfg RouteConfig) {
	auth := api.Group("/auth")
	auth.Post("/login", cfg.AuthHandler.Login)
	auth.Post("/forgot-password", cfg.AuthHandler.ForgotPassword)
	auth.Post("/reset-password/:id", cfg.AuthMiddleware.RequireAuth(), cfg.PermissionMiddleware.Require("USER_UPDATE"), cfg.AuthHandler.AdminResetPassword)
	auth.Post("/change-password", cfg.AuthMiddleware.RequireAuth(), cfg.AuthHandler.ChangePassword)
}

// USER
func UserRoutes(api fiber.Router, cfg RouteConfig) {
	users := api.Group("/users", cfg.AuthMiddleware.RequireAuth())
	users.Get("/", cfg.PermissionMiddleware.Require("USER_VIEW"), cfg.UserHandler.List)
	users.Get("/:id", cfg.PermissionMiddleware.Require("USER_VIEW"), cfg.UserHandler.GetByID)
	users.Post("/create", cfg.PermissionMiddleware.Require("USER_CREATE"), cfg.UserHandler.Create)
	users.Put("/:id", cfg.PermissionMiddleware.Require("USER_UPDATE"), cfg.UserHandler.Update)
	users.Delete("/:id", cfg.PermissionMiddleware.Require("USER_DELETE"), cfg.UserHandler.Delete)
	// USER STATUS
	users.Post("/lock/:id", cfg.PermissionMiddleware.Require("USER_UPDATE"), cfg.UserHandler.Lock)
	users.Post("/unlock/:id", cfg.PermissionMiddleware.Require("USER_UPDATE"), cfg.UserHandler.Unlock)
	// USER PERMISSION OVERRIDES (gộp trong UserHandler)
	users.Get("/overrides/:id", cfg.PermissionMiddleware.Require("USER_PERMISSION_VIEW"), cfg.UserHandler.GetOverrides)
	users.Post("/overrides/assign", cfg.PermissionMiddleware.Require("USER_PERMISSION_ASSIGN"), cfg.UserHandler.AssignOverride)
	users.Delete("/overrides/:id", cfg.PermissionMiddleware.Require("USER_PERMISSION_REMOVE"), cfg.UserHandler.RemoveOverride)
	// Change role user
	users.Put("/change-role/:id", cfg.PermissionMiddleware.Require("USER_UPDATE"), cfg.UserHandler.UpdateRole)
}

// ROLE
func RoleRoutes(api fiber.Router, cfg RouteConfig) {
	roles := api.Group("/roles", cfg.AuthMiddleware.RequireAuth())
	roles.Get("/", cfg.PermissionMiddleware.Require("ROLE_VIEW"), cfg.RoleHandler.List)
	roles.Get("/:id", cfg.PermissionMiddleware.Require("ROLE_VIEW"), cfg.RoleHandler.GetByID)
	roles.Post("/create", cfg.PermissionMiddleware.Require("ROLE_CREATE"), cfg.RoleHandler.Create)
	roles.Put("/:id", cfg.PermissionMiddleware.Require("ROLE_UPDATE"), cfg.RoleHandler.Update)
	roles.Delete("/:id", cfg.PermissionMiddleware.Require("ROLE_DELETE"), cfg.RoleHandler.Delete)
	// ROLE PERMISSIONS (gộp trong RoleHandler)
	roles.Get("/permissions/:id", cfg.PermissionMiddleware.Require("ROLE_PERMISSION_VIEW"), cfg.RoleHandler.GetPermissions)
	roles.Post("/permissions/assign", cfg.PermissionMiddleware.Require("ROLE_PERMISSION_ASSIGN"), cfg.RoleHandler.AssignPermission)
	roles.Delete("/:id/permissions/:permissionId", cfg.PermissionMiddleware.Require("ROLE_PERMISSION_REMOVE"),
		cfg.RoleHandler.RemovePermission)
}

// PERMISSION
func PermissionRoutes(api fiber.Router, cfg RouteConfig) {
	permissions := api.Group("/permissions", cfg.AuthMiddleware.RequireAuth())
	permissions.Get("/", cfg.PermissionMiddleware.Require("PERMISSION_VIEW"), cfg.PermissionHandler.List)
	permissions.Get("/grouped", cfg.PermissionMiddleware.Require("PERMISSION_VIEW"), cfg.PermissionHandler.Grouped)
	permissions.Get("/:id", cfg.PermissionMiddleware.Require("PERMISSION_VIEW"), cfg.PermissionHandler.GetByID)
	permissions.Get("/feature/:featureCode", cfg.PermissionMiddleware.Require("PERMISSION_VIEW"), cfg.PermissionHandler.GetByFeatureCode)
}

// AUDIT
func AuditRoutes(api fiber.Router, cfg RouteConfig) {
	audits := api.Group("/audits", cfg.AuthMiddleware.RequireAuth())
	audits.Get("/", cfg.PermissionMiddleware.Require("AUDIT_VIEW"), cfg.AuditHandler.ListWithFilter)
	audits.Get("/all", cfg.PermissionMiddleware.Require("AUDIT_VIEW"), cfg.AuditHandler.GetAll,)
	audits.Get("/:id", cfg.PermissionMiddleware.Require("AUDIT_VIEW"), cfg.AuditHandler.GetByID)
	audits.Get("/actor/:actorId", cfg.PermissionMiddleware.Require("AUDIT_VIEW"), cfg.AuditHandler.ListByActor)
	audits.Get("/entity", cfg.PermissionMiddleware.Require("AUDIT_VIEW"), cfg.AuditHandler.ListByEntity)
	audits.Get("/tenant", cfg.PermissionMiddleware.Require("AUDIT_VIEW"), cfg.AuditHandler.ListByTenant)
}
