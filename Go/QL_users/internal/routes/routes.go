package routes

import (
	"github.com/gofiber/fiber/v2"

	"go-user-management/internal/handler"
	"go-user-management/internal/middleware"
)

type RouteConfig struct {
	AuthHandler           *handler.AuthHandler
	UserHandler           *handler.UserHandler
	RoleHandler           *handler.RoleHandler
	PermissionHandler     *handler.PermissionHandler
	RolePermissionHandler *handler.RolePermissionHandler
	UserOverrideHandler   *handler.UserPermissionOverrideHandler
	PasswordResetHandler  *handler.PasswordResetHandler
	AuditHandler          *handler.AuditHandler

	AuthMiddleware       *middleware.AuthMiddleware
	PermissionMiddleware *middleware.PermissionMiddleware
}

func RegisterRoutes(app *fiber.App, cfg RouteConfig,) {
	api := app.Group("/api")
	registerAuthRoutes(api, cfg,)
	registerUserRoutes(api, cfg,)
	registerRoleRoutes(api, cfg,)
	registerPermissionRoutes(api, cfg,)
	registerRolePermissionRoutes(api, cfg,)
	registerUserOverrideRoutes(api, cfg,)
	registerPasswordResetRoutes(api, cfg,)
	registerAuditRoutes(api, cfg,)
}
func registerAuthRoutes(api fiber.Router, cfg RouteConfig,) {
	auth := api.Group("/auth")
	auth.Post("/login", cfg.AuthHandler.Login,)
}

func registerPasswordResetRoutes(api fiber.Router, cfg RouteConfig,) {
	reset := api.Group("/password-resets")
	reset.Post("/forgot-password", cfg.PasswordResetHandler.ForgotPassword,)
	reset.Post("/reset-password", cfg.PasswordResetHandler.ResetPassword,)
}

func registerUserRoutes(api fiber.Router, cfg RouteConfig,) {
	users := api.Group("/users", cfg.AuthMiddleware.RequireAuth(),)
	users.Get("/",cfg.PermissionMiddleware.Require("USER_VIEW",),cfg.UserHandler.ListUsers,)
	users.Get("/:id",cfg.PermissionMiddleware.Require("USER_VIEW",), middleware.DataScope("USER_VIEW",), cfg.UserHandler.GetByID,)
	users.Post("/",cfg.PermissionMiddleware.Require("USER_CREATE"), cfg.UserHandler.CreateUser,)
	users.Put("/:id",cfg.PermissionMiddleware.Require("USER_UPDATE",), middleware.DataScope("USER_UPDATE"), cfg.UserHandler.UpdateUser,)
	users.Delete("/:id", cfg.PermissionMiddleware.Require("USER_DELETE",), middleware.DataScope("USER_DELETE",), cfg.UserHandler.DeleteUser,)
}

func registerRoleRoutes(api fiber.Router, cfg RouteConfig,) {
	roles := api.Group("/roles", cfg.AuthMiddleware.RequireAuth(),)
	roles.Get("/", cfg.PermissionMiddleware.Require("ROLE_VIEW",),cfg.RoleHandler.ListRoles,)
	roles.Get("/:id", cfg.PermissionMiddleware.Require("ROLE_VIEW",),cfg.RoleHandler.GetRoleByID,)
	roles.Post("/", cfg.PermissionMiddleware.Require("ROLE_CREATE",), cfg.RoleHandler.CreateRole,)
	roles.Put("/:id", cfg.PermissionMiddleware.Require("ROLE_UPDATE",), cfg.RoleHandler.UpdateRole,)
	roles.Delete("/:id", cfg.PermissionMiddleware.Require("ROLE_DELETE",), cfg.RoleHandler.DeleteRole,)
}

func registerPermissionRoutes(api fiber.Router,cfg RouteConfig,) {
	permissions := api.Group("/permissions",cfg.AuthMiddleware.RequireAuth(),)
	permissions.Get("/", cfg.PermissionMiddleware.Require("PERMISSION_VIEW",), cfg.PermissionHandler.GetAllPermissions,)
	permissions.Get("/:id", cfg.PermissionMiddleware.Require("PERMISSION_VIEW",), cfg.PermissionHandler.GetPermissionByID,)
	permissions.Get("/feature/:featureCode", cfg.PermissionMiddleware.Require("PERMISSION_VIEW",),cfg.PermissionHandler.GetPermissionByFeatureCode,)
	permissions.Get("/role/:roleId", cfg.PermissionMiddleware.Require("PERMISSION_VIEW",), cfg.PermissionHandler.GetPermissionsByRole,)
	permissions.Get("/role/:roleId/scope", cfg.PermissionMiddleware.Require("PERMISSION_VIEW",),cfg.PermissionHandler.GetPermissionsByRoleWithScope,)
}

func registerRolePermissionRoutes(api fiber.Router, cfg RouteConfig,) {
	rp := api.Group("/role-permissions", cfg.AuthMiddleware.RequireAuth(),)
	rp.Get("/role/:roleId", cfg.PermissionMiddleware.Require("ROLE_PERMISSION_VIEW",), cfg.RolePermissionHandler.GetByRoleID,)
	rp.Post("/assign", cfg.PermissionMiddleware.Require("ROLE_PERMISSION_ASSIGN",), cfg.RolePermissionHandler.AssignPermission,)
	rp.Post("/revoke", cfg.PermissionMiddleware.Require("ROLE_PERMISSION_REVOKE",), cfg.RolePermissionHandler.RevokePermission,)
}

func registerUserOverrideRoutes(api fiber.Router, cfg RouteConfig,) {
	override := api.Group("/user-permission-overrides",cfg.AuthMiddleware.RequireAuth(),)
	override.Get("/user/:userId", cfg.PermissionMiddleware.Require("USER_OVERRIDE_VIEW",), cfg.UserOverrideHandler.GetByUserID,)
	override.Post("/grant", cfg.PermissionMiddleware.Require("USER_OVERRIDE_GRANT",), cfg.UserOverrideHandler.GrantOverride,)
	override.Post("/revoke", cfg.PermissionMiddleware.Require("USER_OVERRIDE_REVOKE",), cfg.UserOverrideHandler.RevokeOverride,)
}

func registerAuditRoutes(api fiber.Router, cfg RouteConfig,) {
	audit := api.Group("/audit-logs",cfg.AuthMiddleware.RequireAuth(),)
	audit.Get("/", cfg.PermissionMiddleware.Require("AUDIT_VIEW",), cfg.AuditHandler.List,)
	audit.Get("/:id", cfg.PermissionMiddleware.Require("AUDIT_VIEW",), cfg.AuditHandler.GetByID,)
	audit.Post("/", cfg.PermissionMiddleware.Require("AUDIT_CREATE",), cfg.AuditHandler.Create,)
}
