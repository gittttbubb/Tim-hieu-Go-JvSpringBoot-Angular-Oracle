package authz

import "strings"

type PermissionGrant struct {
	PermissionCode string
	Granted        bool
	DataScope      string
}

func BuildPermissionKey(featureGroup, action string) string {
	return strings.ToUpper(featureGroup) + "." + strings.ToUpper(action)
}

type PermissionResolver interface {
	HasPermission(permissionCode string) bool
	GetDataScope(permissionCode string) string
	GetPermission(permissionCode string) (*PermissionGrant, bool)
	GetAllPermissions() []PermissionGrant
}

type permissionResolver struct {
	permissions map[string]PermissionGrant
}

func NewPermissionResolver(grants []PermissionGrant) PermissionResolver {
	permissions := make(map[string]PermissionGrant)
	for _, grant := range grants {
		permissions[grant.PermissionCode] = grant
	}
	return &permissionResolver{permissions: permissions}
}

func (r *permissionResolver) HasPermission(permissionCode string) bool {
	permission, exists := r.permissions[permissionCode]
	if !exists {
		return false
	}
	return permission.Granted
}

func (r *permissionResolver) GetDataScope(permissionCode string) string {
	permission, exists := r.permissions[permissionCode]
	if !exists {
		return ""
	}
	if !permission.Granted {
		return ""
	}
	return permission.DataScope
}

func (r *permissionResolver) GetPermission(permissionCode string) (*PermissionGrant, bool) {
	permission, exists := r.permissions[permissionCode]
	if !exists {
		return nil, false
	}
	return &permission, true
}

func (r *permissionResolver) GetAllPermissions() []PermissionGrant {
	result := make([]PermissionGrant, 0, len(r.permissions))
	for _, permission := range r.permissions {
		result = append(result, permission)
	}
	return result
}
