package dto

type AssignRolePermissionRequest struct {
	RoleID       string `json:"roleId" validate:"required,uuid"`
	PermissionID string `json:"permissionId" validate:"required"`
	Granted      bool   `json:"granted"`
	DataScope    string `json:"dataScope" validate:"required,oneof=OWN TEAM ALL"`
}

type RolePermissionResponse struct {
	RoleID       string `json:"roleId"`
	PermissionID string `json:"permissionId"`
	Granted      bool   `json:"granted"`
	DataScope    string `json:"dataScope"`
}