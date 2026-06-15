package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	_"go-user-management/internal/model"
	"go-user-management/internal/utils"
)

type AuthMiddleware struct {
	jwtSecret string
}

func NewAuthMiddleware(jwtSecret string,) *AuthMiddleware {
	return &AuthMiddleware{
		jwtSecret: jwtSecret,
	}
}

func (m *AuthMiddleware) RequireAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return utils.Unauthorized(
				c,
				"authorization header is required",
			)
		}
		if !strings.HasPrefix(authHeader, "Bearer ",) {
			return utils.Unauthorized(
				c,
				"invalid authorization header",
			)
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ",)
		claims, err := utils.ParseJWT(tokenString, m.jwtSecret,)
		if err != nil {
			return utils.Unauthorized(
				c,
				"invalid or expired token",
			)
		}
		c.Locals("claims", claims,)
		return c.Next()
	}
}