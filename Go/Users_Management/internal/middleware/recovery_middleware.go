package middleware

import (
	"fmt"
	"runtime/debug"

	"go-rbac-system/internal/constants"
	"go-rbac-system/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type RecoveryMiddleware interface {
	Handle() fiber.Handler
}

type recoveryMiddleware struct{}

func NewRecoveryMiddleware() RecoveryMiddleware {
	return &recoveryMiddleware{}
}
// Nếu có panic xảy ra ở bất kỳ đâu khi request tới sẽ log ra http 500
func (m *recoveryMiddleware) Handle() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		defer func() {
			if r := recover(); r != nil {

				// Extract request info
				method := c.Method()
				path := c.Path()
				ip := c.IP()

				// Extract user context (optional)
				var userID string
				var tenantID string

				if claims := c.Locals(constants.ContextClaims); claims != nil {
					// safe type assertion (avoid panic inside recovery)
					if v, ok := claims.(interface {
						GetUserID() string
						GetTenantID() string
					}); ok {
						userID = v.GetUserID()
						tenantID = v.GetTenantID()
					}
				}

				// Stack trace
				stack := string(debug.Stack())

				// Log (basic production logging)
				fmt.Printf("\n[RECOVERY PANIC]\n")
				fmt.Printf("error: %v\n", r)
				fmt.Printf("method: %s | path: %s | ip: %s\n", method, path, ip)
				fmt.Printf("user_id: %s | tenant_id: %s\n", userID, tenantID)
				fmt.Printf("stack:\n%s\n", stack)

				// Response
				_ = response.Error(
					c,
					fiber.StatusInternalServerError,
					"internal server error",
				)

				err = nil
			}
		}()

		return c.Next()
	}
}