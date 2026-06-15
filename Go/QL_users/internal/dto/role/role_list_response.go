package dto

type RoleListResponse struct {
	Items []RoleResponse `json:"items"`
	Total int            `json:"total"`
}