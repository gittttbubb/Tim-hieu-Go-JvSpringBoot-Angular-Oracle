package dto

type UserPermissionOverrideRequest struct {
	UserID       string `json:"userId" validate:"required,uuid"`
	PermissionID string `json:"permissionId" validate:"required"`
	Granted      bool   `json:"granted"`
	Reason string `json:"reason" validate:"max=500"`
	DataScope string `json:"dataScope" validate:"required,oneof=OWN TEAM ALL"`
}

type UserPermissionOverrideResponse struct {
	ID           string `json:"id"`
	UserID       string `json:"userId"`
	PermissionID string `json:"permissionId"`
	Granted      bool   `json:"granted"`
	Reason       string `json:"reason"`
	CreatedBy    string `json:"createdBy"`
	DataScope    string `json:"dataScope"`
}