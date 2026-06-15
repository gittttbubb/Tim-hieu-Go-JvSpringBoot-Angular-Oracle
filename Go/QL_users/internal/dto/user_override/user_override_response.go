package dto

type UserPermissionOverrideResponse struct {
	ID           string `json:"id"`
	UserID       string `json:"userId"`
	PermissionID string `json:"permissionId"`
	Granted      bool   `json:"granted"`
	Reason       string `json:"reason"`
	CreatedBy    string `json:"createdBy"`
	DataScope    string `json:"dataScope"`
}