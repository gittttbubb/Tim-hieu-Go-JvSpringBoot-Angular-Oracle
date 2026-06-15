package dto

type PermissionListResponse struct {
	Items []PermissionResponse `json:"items"`
	Total int                  `json:"total"`
}