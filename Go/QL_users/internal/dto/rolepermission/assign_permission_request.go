package dto

type AssignPermissionRequest struct {
	RoleID       string `json:"roleId"`
	PermissionID string `json:"permissionId"`
	DataScope    string `json:"dataScope"`
}