package dto

type RolePermissionListResponse struct {
	Items []RolePermissionResponse `json:"items"`
	Total int                      `json:"total"`
}