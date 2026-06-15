package dto

type UserPermissionOverrideListResponse struct {
	Items []UserPermissionOverrideResponse `json:"items"`
	Total int                              `json:"total"`
}