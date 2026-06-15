package dto

type UpdateUserRequest struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	RoleID   string `json:"roleId"`
	Status   string `json:"status"`
}