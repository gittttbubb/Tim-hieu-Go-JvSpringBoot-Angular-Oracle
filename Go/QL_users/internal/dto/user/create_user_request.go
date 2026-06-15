package dto

type CreateUserRequest struct {
	TenantID            string `json:"tenantId"`
	FullName            string `json:"fullName"`
	Username            string `json:"username"`
	Email               string `json:"email"`
	Phone               string `json:"phone"`
	Password            string `json:"password"`
	RoleID              string `json:"roleId"`
	MustChangePassword  bool   `json:"mustChangePassword"`
}