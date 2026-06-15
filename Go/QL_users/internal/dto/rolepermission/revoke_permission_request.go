package dto

type RevokePermissionRequest struct {
	RoleID       string `json:"roleId"`
	PermissionID string `json:"permissionId"`
}