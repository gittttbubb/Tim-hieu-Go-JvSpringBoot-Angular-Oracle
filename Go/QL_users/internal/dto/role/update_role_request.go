package dto

type UpdateRoleRequest struct {
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
}