package dto

type CreateUserRequest struct {
	TenantID string `json:"tenantId" validate:"required,uuid"`
	FullName string `json:"fullName" validate:"required,max=100"`
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email,max=200"`
	Phone    string `json:"phone" validate:"required,max=20"`
	RoleID   string `json:"roleId" validate:"required,uuid"`
	// Password string `json:"password" validate:"required,min=8,max=100"`
}

type UpdateUserRequest struct {
	FullName string `json:"fullName" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email,max=200"`
	Phone    string `json:"phone" validate:"required,max=20"`
	RoleID   string `json:"roleId" validate:"required,uuid"`
	Status   string `json:"status" validate:"required,oneof=ACTIVE LOCKED PENDING_PASSWORD_CHANGE"`
}

type UserListResponse struct {
	ID                 string `json:"id"`
	TenantID           string `json:"tenantId"`
	FullName           string `json:"fullName"`
	Username           string `json:"username"`
	Email              string `json:"email"`
	Phone              string `json:"phone"`
	RoleID             string `json:"roleId"`
	Status             string `json:"status"`
	MustChangePassword bool   `json:"mustChangePassword"`
}

type UserDetailResponse struct {
	ID                 string `json:"id"`
	TenantID           string `json:"tenantId"`
	FullName           string `json:"fullName"`
	Username           string `json:"username"`
	Email              string `json:"email"`
	Phone              string `json:"phone"`
	RoleID             string `json:"roleId"`
	Status             string `json:"status"`
	MustChangePassword bool   `json:"mustChangePassword"`
	CreatedAt          string `json:"createdAt"`
	UpdatedAt          string `json:"updatedAt"`
}

type UpdateUserRoleRequest struct {
    RoleID string `json:"roleId" validate:"required,uuid"`
}