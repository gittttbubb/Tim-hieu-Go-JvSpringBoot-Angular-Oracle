package middleware

import (
	"strings"

	"go-rbac-system/internal/config"
	"go-rbac-system/internal/constants"
	"go-rbac-system/pkg/response"
	"go-rbac-system/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware interface {
	RequireAuth() fiber.Handler
}

type authMiddleware struct {
	cfg *config.Config
}

func NewAuthMiddleware(
	cfg *config.Config,
) AuthMiddleware {
	return &authMiddleware{
		cfg: cfg,
	}
}

func (m *authMiddleware) RequireAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Lấy token từ http header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return response.Error(
				c,
				fiber.StatusUnauthorized,
				"authorization header is required",
			)
		}
		// tách chuỗi làm 2 phần dựa vào khoảng trắng " "
		parts := strings.SplitN(authHeader, " ", 2,)
		if len(parts) != 2 {
			return response.Error(
				c,
				fiber.StatusUnauthorized,
				"invalid authorization format",
			)
		}
		// Kiểm tra phần đầu tiên có phải chữ Bearer
		if !strings.EqualFold(parts[0], "Bearer") {
			return response.Error(
				c,
				fiber.StatusUnauthorized,
				"invalid authorization format",
			)
		}
		tokenString := strings.TrimSpace(parts[1])
		if tokenString == "" {
			return response.Error(
				c,
				fiber.StatusUnauthorized,
				"token is required",
			)
		}
		claims, err := utils.ParseToken(
			tokenString,
			m.cfg.JWT.Secret,
		)
		if err != nil {
			return response.Error(
				c,
				fiber.StatusUnauthorized,
				"invalid or expired token",
			)
		}

		c.Locals(constants.ContextClaims, claims)
		c.Locals(constants.ContextUserID, claims.UserID)
		c.Locals(constants.ContextRoleIDs, claims.RoleID)

		return c.Next()
	}
}