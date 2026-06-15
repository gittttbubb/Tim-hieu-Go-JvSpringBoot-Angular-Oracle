package model

type RolePermission struct {
	RoleID       string `json:"roleId"`
	PermissionID string `json:"permissionId"`
	Granted      bool   `json:"granted"`
	DataScope    string `json:"dataScope"`
}