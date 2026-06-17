package model

type RolePermission struct {
	RoleID       string `json:"roleId" db:"role_id"`
	PermissionID string `json:"permissionId" db:"permission_id"`
	Granted      bool   `json:"granted" db:"granted"`
	DataScope    string `json:"dataScope" db:"data_scope"`
}