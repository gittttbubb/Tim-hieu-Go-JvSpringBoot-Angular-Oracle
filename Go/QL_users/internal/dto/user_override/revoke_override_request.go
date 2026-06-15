package dto

type RevokeOverrideRequest struct {
	UserID       string `json:"userId"`
	PermissionID string `json:"permissionId"`
}