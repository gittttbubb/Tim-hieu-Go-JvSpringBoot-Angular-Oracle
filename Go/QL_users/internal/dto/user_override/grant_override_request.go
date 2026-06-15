package dto

type GrantOverrideRequest struct {
	UserID       string `json:"userId"`
	PermissionID string `json:"permissionId"`
	Granted      bool   `json:"granted"`
	Reason       string `json:"reason"`
	DataScope    string `json:"dataScope"`
}