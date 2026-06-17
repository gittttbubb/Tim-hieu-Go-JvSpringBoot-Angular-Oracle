package middleware

import (

	"go-rbac-system/internal/constants"
	"go-rbac-system/internal/service"
	"go-rbac-system/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type AuditMiddleware interface {
	Log(action string) fiber.Handler
}

type auditMiddleware struct {
	auditService service.AuditService
}

func NewAuditMiddleware(
	auditService service.AuditService,
) AuditMiddleware {
	return &auditMiddleware{
		auditService: auditService,
	}
}

func (m *auditMiddleware) Log(action string) fiber.Handler {
	return func(c *fiber.Ctx) error {

		// =====================
		// Capture request info ONLY
		// =====================
		method := c.Method()
		path := c.Path()
		ip := c.IP()

		// =====================
		// Extract claims ONLY
		// =====================
		var userID string
		var tenantID string
		var actor string

		if v := c.Locals(constants.ContextClaims); v != nil {
			if claims, ok := v.(*utils.Claims); ok {
				userID = claims.UserID
				tenantID = claims.TenantID
				actor = claims.Username
			}
		}

		// =====================
		// Execute handler FIRST
		// =====================
		err := c.Next()

		// =====================
		// Build context payload ONLY
		// (NO business logic here)
		// =====================
		auditCtx := &service.AuditContext{
			UserID:   userID,
			TenantID: tenantID,
			Actor:    actor,
			Method:   method,
			Path:     path,
			IP:       ip,
		}

		// =====================
		// Delegate ALL logic to service
		// =====================
		_ = m.auditService.Log(action, auditCtx)

		return err
	}
}