package middleware

import (
	"github.com/gofiber/fiber/v2"

	"go-user-management/internal/model"
	"go-user-management/internal/utils"
)

type PermissionMiddleware struct {
}

func NewPermissionMiddleware() *PermissionMiddleware {
	return &PermissionMiddleware{}
}	

func (m *PermissionMiddleware) Require(featureCode string,) fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := GetClaims(c.Locals("claims"),)
		if claims == nil {
			return utils.Unauthorized(
				c,
				"unauthorized",
			)
		}
		for _, permission := range claims.Permissions {
			if permission.FeatureCode == featureCode {
				return c.Next()
			}
		}
		return utils.Forbidden(
			c,
			"permission denied",
		)
	}
}

func (m *PermissionMiddleware) GetScope(claims *model.JWTClaims,featureCode string,) string {
	for _, permission := range claims.Permissions {
		if permission.FeatureCode == featureCode {
			return permission.DataScope
		}
	}
	return ""
}