package bootstrap

import (
	"go-rbac-system/internal/middleware"
	"go-rbac-system/internal/routes"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2"
)

func NewApp() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "go-rbac-system",
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:4200",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin,Content-Type,Accept,Authorization,Accept-Language,X-Lang",
		AllowCredentials: true,
	}))

	// Register i18n middleware globally
	app.Use(middleware.I18nMiddleware())

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "UP",
		})
	})
	return app
}

func BuildRouteConfig(c *Container) routes.RouteConfig {
	return routes.RouteConfig{
		AuthHandler:       c.AuthHandler,
		UserHandler:       c.UserHandler,
		RoleHandler:       c.RoleHandler,
		PermissionHandler: c.PermissionHandler,
		AuditHandler:      c.AuditHandler,
		AuthMiddleware:       c.AuthMiddleware,
		PermissionMiddleware: c.PermissionMiddleware,
	}
}