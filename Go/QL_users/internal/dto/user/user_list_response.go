package dto

type UserListResponse struct {
	Items []UserResponse `json:"items"`
	Total int            `json:"total"`
}