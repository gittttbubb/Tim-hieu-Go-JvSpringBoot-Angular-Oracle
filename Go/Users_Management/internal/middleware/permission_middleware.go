package middleware

import (
	"go-rbac-system/internal/authz"
	"go-rbac-system/internal/constants"
	"go-rbac-system/internal/model"
	"go-rbac-system/internal/repository"
	"go-rbac-system/internal/service"
	"go-rbac-system/pkg/response"
	"go-rbac-system/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type PermissionMiddleware interface {
	Require(permissionCode string) fiber.Handler
	RequireAny(permissionCodes ...string) fiber.Handler
}

type permissionMiddleware struct {
	authService        service.AuthService
	permissionRepo     repository.PermissionRepository
}

func NewPermissionMiddleware(authService service.AuthService, permissionRepo repository.PermissionRepository,) PermissionMiddleware {
	return &permissionMiddleware{
		authService:    authService,
		permissionRepo: permissionRepo,
	}
}

func (m *permissionMiddleware) Require(permissionCode string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 1. Lấy thông tin từ context
		claimsValue := c.Locals(constants.ContextClaims)
		if claimsValue == nil {
			return response.Error(
				c,
				fiber.StatusUnauthorized,
				"unauthorized",
			)
		}
		// Ép kiểu Claims
		claims, ok := claimsValue.(*utils.Claims)
		if !ok {
			return response.Error(
				c,
				fiber.StatusUnauthorized,
				"invalid auth context",
			)
		}
		// 2. Load quyền từ DB
		rolePermissions, userOverrides, err := m.authService.LoadAuthorizationData(claims.UserID, claims.RoleID,)
		if err != nil {
			return response.Error(
				c,
				fiber.StatusInternalServerError,
				"failed to load authorization data",
			)
		}
		// 3. Load permission catalog (ID -> Code mapping)
		permissions, err := m.permissionRepo.ListAll()
		if err != nil {
			return response.Error(
				c,
				fiber.StatusInternalServerError,
				"failed to load permissions",
			)
		}
		permissionCodeMap := buildPermissionCodeMap(permissions)
		// 4. Build Permission Grants
		grants := buildPermissionGrants(
			rolePermissions,
			userOverrides,
			permissionCodeMap,
		)
		// 5. Create resolver (per request)
		resolver := authz.NewPermissionResolver(grants)
		// 6. Check permission
		if !resolver.HasPermission(permissionCode) {
			return response.Error(
				c,
				fiber.StatusForbidden,
				"permission denied",
			)
		}
		// 7. Resolve data scope
		scope := resolver.GetDataScope(permissionCode)
		// 8. Store in context
		c.Locals(constants.ContextPermissionGrants, grants)
		c.Locals(constants.ContextDataScope, scope)
		c.Locals("permission_resolver", resolver)
		// 9. Continue
		return c.Next()
	}
}
func (m *permissionMiddleware) RequireAny(permissionCodes ...string) fiber.Handler {
    return func(c *fiber.Ctx) error {
        claimsValue := c.Locals(constants.ContextClaims)
        if claimsValue == nil {
            return response.Error(
                c,
                fiber.StatusUnauthorized,
                "unauthorized",
            )
        }
        claims, ok := claimsValue.(*utils.Claims)
        if !ok {
            return response.Error(
                c,
                fiber.StatusUnauthorized,
                "invalid auth context",
            )
        }
        rolePermissions, userOverrides, err := m.authService.LoadAuthorizationData(
            claims.UserID,
            claims.RoleID,
        )
        if err != nil {
            return response.Error(
                c,
                fiber.StatusInternalServerError,
                "failed to load authorization data",
            )
        }
        permissions, err := m.permissionRepo.ListAll()
        if err != nil {
            return response.Error(
                c,
                fiber.StatusInternalServerError,
                "failed to load permissions",
            )
        }
        permissionCodeMap := buildPermissionCodeMap(permissions)
        grants := buildPermissionGrants(
            rolePermissions,
            userOverrides,
            permissionCodeMap,
        )
        resolver := authz.NewPermissionResolver(grants)
        for _, code := range permissionCodes {
            if resolver.HasPermission(code) {
                c.Locals(constants.ContextPermissionGrants, grants)
                return c.Next()
            }
        }
        return response.Error(
            c,
            fiber.StatusForbidden,
            "permission denied",
        )
    }
}
// Helpers
// tạo map[string]string. Kết quả VD:{
//     "p1": "USER_VIEW",
//     "p2": "USER_CREATE",
//     "p3": "USER_DELETE",
// }
func buildPermissionCodeMap(permissions []model.Permission,) map[string]string {
	m := make(map[string]string, len(permissions))
	for _, p := range permissions {
		m[p.ID] = p.FeatureCode
	}
	return m
}

func buildPermissionGrants(rolePermissions []model.RolePermission, userOverrides []model.UserPermissionOverride, 
	permissionCodeMap map[string]string) []authz.PermissionGrant {
	grantMap := make(map[string]authz.PermissionGrant)
	// 1. Role permissions first
	for _, rp := range rolePermissions {
		code := permissionCodeMap[rp.PermissionID]
		if code == "" {
			continue
		}
		grantMap[code] = authz.PermissionGrant{
			PermissionCode: code,
			Granted:        rp.Granted,
			DataScope:      rp.DataScope,
		}
	}
	// 2. User overrides (override role)
	for _, up := range userOverrides {
		code := permissionCodeMap[up.PermissionID]
		if code == "" {
			continue
		}
		grantMap[code] = authz.PermissionGrant{
			PermissionCode: code,
			Granted:        up.Granted,
			DataScope:      up.DataScope,
		}
	}
	// 3. Convert map → slice
	grants := make([]authz.PermissionGrant, 0, len(grantMap))
	for _, g := range grantMap {
		grants = append(grants, g)
	}
	return grants
}