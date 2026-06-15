package middleware

import (
	"go-user-management/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func DataScope(featureCode string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := GetClaims(c.Locals("claims"),)
		if claims == nil {
			return utils.Unauthorized(
				c,
				"invalid token",
			)
		}
		var scope string
		for _, permission := range claims.Permissions {
			if permission.FeatureCode == featureCode {
				scope = permission.DataScope
				break
			}
		}
		if scope == "" {
			return utils.Forbidden(
				c,
				"permission denied",
			)
		}
		if scope == "ALL" {
			return c.Next()
		}
		if scope == "OWN" {
			targetID := c.Params("id")
			if targetID == "" {
				return utils.Forbidden(
					c,
					"own scope requires resource id",
				)
			}
			if targetID != claims.UserID {
				return utils.Forbidden(
					c,
					"access denied",
				)
			}
		}
		return c.Next()
	}
}